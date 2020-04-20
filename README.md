```text
     ____  __
    / __ \/ /_______      __
   / /_/ / __/ ___/ | /| / /
  / ____/ /_/ /   | |/ |/ /
 /_/    \__/_/    |__/|__(_)

   Sasha P. <dev@ptrw.nl>
----------------------------
```

> tlsify: a TLS termination proxy

Build using `go build -o tlsify ./src`.

Run using `tlsify <type> <address> <type> <address> <certificate> <key>`, where
`<type>` indicates a transport type, `<address>` specifies a connection address
and `<certificate>` and `<key>` are paths to an X.509 certificate and
its key.

```ebnf
type = ( "tc" | "ud" ) , "p" , [ "4" | "6" ]
     | "unix" , [ "gram" | "packet" ] ;
```

Converting an IPv4 HTTP endpoint to an IPv6 HTTPS endpoint can be done using
`tlsify tcp4 :80 tcp6 :443 ./server.cert ./server.key`. Make sure the
unencrypted endpoint is only part of a private network if sensitive data is
transferred.

This repository contains a self-signed certificate for testing, do not use this
certificate for any other purpose. Feel free to contact me if you have trouble
obtaining a certificate or configuring the software.

```text
tlsify: a TLS termination proxy
Copyright (C) 2020  Sasha P. <dev@ptrw.nl>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```
