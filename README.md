## A tiny, readable HTTP/1.1 server that:

- Opens a TCP socket and listens on a port.
- Accepts incoming connections and spawns concurrent handlers.
- Reads the raw request line (e.g. GET /hello HTTP/1.1).
- Parses headers RFC-style into a map[string][]string (case-insensitive, supports multiple values).
- Validates required headers (e.g. Host for HTTP/1.1).
- Builds a well-formed HTTP response with correct CRLF (\r\n) and Content-Length.

**This server is educational. It intentionally focuses on clarity rather than production-ready robustness.**

## Run locally

```
git clone https://github.com/Sai-Myo-Myat/toy_server.git //https

cd toy_server

go mod tidy

go run .
```
