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

Build using `go build` or install using `go get github.com/tlsify/tlsify`.

Run using `tlsify <type> <address> <type> <address> <certificate> <key>`, where
`<type>` indicates a transport type, `<address>` specifies a connection address
and `<certificate>` and `<key>` are paths to an X.509 certificate and its key.

```ebnf
type = ( "tc" | "ud" ) , "p" , [ "4" | "6" ]
     | "unix" , [ "gram" | "packet" ] ;
```

Converting an IPv4 HTTP endpoint to an IPv6 HTTPS endpoint can be done using
`tlsify tcp4 :80 tcp6 :443 test.cert test.key`. Make sure the unencrypted
endpoint is only part of a private network if sensitive data is transferred.
This repository contains a self-signed certificate for testing, do not use this
certificate for any other purpose. Feel free to contact me if you have trouble
obtaining a certificate or configuring the software.