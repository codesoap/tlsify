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
	"io"
	"log"
	"net"
	"sync"
)

func abort(err error) {
	problem(err, log.Fatal)
}

func problem(err error, hndl func(...interface{})) bool {

	if err != nil {
		hndl(err)
		return true
	}

	return false
}

func stream(dst net.Conn, src net.Conn, grp *sync.WaitGroup) {
	defer grp.Done()
	_, err := io.Copy(dst, src)
	warn(err)
}

func warn(err error) bool {
	return problem(err, log.Print)
}
