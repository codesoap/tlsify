// tlsify: a TLS termination proxy
// Copyright (C) 2020  Sasha P. <dev@ptrw.nl>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	banner = `
     ____  __
    / __ \/ /_______      __
   / /_/ / __/ ___/ | /| / /
  / ____/ /_/ /   | |/ |/ /
 /_/    \__/_/    |__/|__(_)

   Sasha P. <dev@ptrw.nl>
----------------------------
`
	usage = `
Usage: %s <type> <address> <type> <address> <certificate> <key>
`
	certCacheDuration = time.Minute
)

var (
	cert                  *tls.Certificate
	certCachedAt          time.Time
	cachedCertFileModTime time.Time
	cachedKeyFileModTime  time.Time
)

func main() {
	fmt.Print(banner)

	if len(os.Args) != 7 {
		fmt.Printf(usage, os.Args[0])
		os.Exit(1)
	}

	srvr, err := tls.Listen(os.Args[3], os.Args[4], &tls.Config{
		GetCertificate: getCertificate,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer srvr.Close()
	log.Print("tlsify is running")

	for {
		clnt, err := srvr.Accept()

		go func() {
			if err != nil {
				log.Print(err)
				return
			}

			defer clnt.Close()
			srvc, err := net.Dial(os.Args[1], os.Args[2])

			if err != nil {
				log.Print(err)
				return
			}

			defer srvc.Close()
			rslt := make(chan error, 1)
			go stream(clnt, srvc, rslt)
			go stream(srvc, clnt, rslt)

			if err := <-rslt; err != nil {
				log.Print(err)
			}
		}()
	}
}

func getCertificate(_ *tls.ClientHelloInfo) (*tls.Certificate, error) {
	if cert == nil || time.Now().After(certCachedAt.Add(certCacheDuration)) {
		var certStat, keyStat os.FileInfo
		certStat, err := os.Stat(os.Args[5])
		if err != nil {
			return nil, fmt.Errorf("could not read FileInfo for certificate: %w", err)
		}
		keyStat, err = os.Stat(os.Args[6])
		if err != nil {
			return nil, fmt.Errorf("could not read FileInfo for key: %w", err)
		}
		if certStat.ModTime() != cachedCertFileModTime ||
			keyStat.ModTime() != cachedKeyFileModTime {
			cachedCertFileModTime = certStat.ModTime()
			cachedKeyFileModTime = keyStat.ModTime()
			log.Print("loading certificate")
			c, err := tls.LoadX509KeyPair(os.Args[5], os.Args[6])
			if err != nil {
				return nil, fmt.Errorf("could not load certificate: %w", err)
			}
			cert = &c
		}
		certCachedAt = time.Now()
	}
	return cert, nil
}

func stream(dst net.Conn, src net.Conn, rslt chan error) {
	_, err := io.Copy(dst, src)
	rslt <- err
}
