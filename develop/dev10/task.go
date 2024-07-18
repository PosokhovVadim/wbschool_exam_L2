package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
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

var (
	timeout time.Duration
)

func tcpConnection(host string, port string) error {
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)

	if err != nil {
		return err
	}
	defer conn.Close()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)
	doneChan := make(chan struct{})


	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		conn.Close()
		close(doneChan)
	}()


	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		close(doneChan)
	}()

	select {
	case <-sigChan:
		return fmt.Errorf("interrupted")
	case <-doneChan:
		fmt.Println("Connection closed, shutting down...")
	}
	return nil
}

func run() error {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout")
	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		return fmt.Errorf("expected host and port, got %v", args)
	}

	host := args[0]
	port := args[1]

	return tcpConnection(host, port)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
