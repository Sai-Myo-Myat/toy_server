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

    fmt.Println("Server running on http://localhost:8080")

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

    method := parts[0]
    path := parts[1]
    version := parts[2]

    // Read headers
    headers := make(map[string]string)
    for {
        line, _ := reader.ReadString('\n')
        line = strings.TrimSpace(line)
        if line == "" {
            break
        }
        h := strings.SplitN(line, ":", 2)
        if len(h) == 2 {
            headers[strings.TrimSpace(h[0])] = strings.TrimSpace(h[1])
        }
    }

    fmt.Println("Method:", method, "Path:", path, "Version:", version)
    fmt.Println("Headers:", headers)

    // Response
    body := fmt.Sprintf("<h1>You requested %s</h1>", path)
    response := "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/html\r\n" +
        fmt.Sprintf("Content-Length: %d\r\n", len(body)) +
        "\r\n" + body

    conn.Write([]byte(response))
}
