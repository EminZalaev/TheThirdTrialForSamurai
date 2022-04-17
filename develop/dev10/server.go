package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8083")
	if err != nil {
		fmt.Println("Cannot start server")
	}

	conn, _ := ln.Accept()

	for {
		mes, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			return
		}

		fmt.Print("Server: ", mes)

		mes = "socket " + mes

		conn.Write([]byte(mes + "\n"))
	}
}
