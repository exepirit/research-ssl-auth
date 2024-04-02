# research-ssl-auth
Client and server-side authentication research.

These features planned and may be implemented:
- Server with TLS authentication.
  - [x] Go
  - [ ] C#/.NET
  - [ ] Python
- [ ] Client with own TLS certificate.
  - [ ] Go
  - [ ] C#/.NET
  - [ ] Python

## Quickstart

Generate CA and server certificate.
```shell
./generate_root_ca.sh
./generate_cert.sh server localhost
```

This will generate two certificate/private key pairs: `root.crt`/`root.key` for local CA, and `server.crt`/`server.key`
for HTTPS server. **CA keys pair should only be used to issue certificates**.

### Start Go server

Server requires certificate/private key pair, which used to establish a TLS-secured connection.

```shell
cd go/
go run ./cmd/server -cert ../server.crt -key ../server.key
```
