package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {

	timeOut := flag.Int("timeout", 10, "Time out flag")
	flag.Parse()
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments.\nUSE : -timeout <timeout> <host> <port>")
		return
	}

	conn, err := net.DialTimeout("tcp", os.Args[3] + ":" + os.Args[4], time.Duration(*timeOut) * time.Second)
	if err != nil {
		time.After(time.Duration(*timeOut) * time.Second)
		fmt.Println("Wrong server ip")
		return
	}

	if conn != nil {
		defer conn.Close()
		fmt.Println("Client Opened")
	}

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err == io.EOF {
				conn.Close()
			}

			fmt.Fprint(conn, text+"\n")
		}
	}()

	for {
		mes, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("Client: " + mes)
	}
}
