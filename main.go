package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer ln.Close()
    fmt.Println("RFC-like Server running on http://localhost:8080")

    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)

    requestLine, _ := reader.ReadString('\n')
    parts := strings.Fields(requestLine)
    if len(parts) < 3 {
        return
    }
    method, path, version := parts[0], parts[1], parts[2]

    headers := make(map[string][]string)
    for {
        line, _ := reader.ReadString('\n')
        line = strings.TrimRight(line, "\r\n")
        if line == "" {
            break
        }
        parts := strings.SplitN(line, ":", 2)
        if len(parts) == 2 {
            key := canonicalHeaderKey(strings.TrimSpace(parts[0]))
            val := strings.TrimSpace(parts[1])
            headers[key] = append(headers[key], val)
        }
    }

    fmt.Println("Method:", method, "Path:", path, "Version:", version)
    fmt.Println("Headers:", headers)

    if _, ok := headers["Host"]; !ok && version == "HTTP/1.1" {
        badReq := "HTTP/1.1 400 Bad Request\r\n" +
            "Content-Length: 0\r\n\r\n"
        conn.Write([]byte(badReq))
        return
    }

    body := fmt.Sprintf("<h1>You requested %s</h1>", path)
    response := "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/html\r\n" +
        fmt.Sprintf("Content-Length: %d\r\n", len(body)) +
        "\r\n" + body
    conn.Write([]byte(response))
}

func canonicalHeaderKey(s string) string {
    s = strings.ToLower(s)
    parts := strings.Split(s, "-")
    for i := range parts {
        if len(parts[i]) > 0 {
            parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
        }
    }
    return strings.Join(parts, "-")
}
