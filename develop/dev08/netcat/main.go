package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	udp := flag.Bool("u", false, "Use UDP (default is TCP)")
	
	flag.Parse()

	protocol := "tcp"
	if *udp {
		protocol = "udp"
	}

	if len(flag.Args()) != 2 {
		fmt.Println("Usage: nc [OPTIONS] HOST PORT")
		os.Exit(1)
	}
	host := flag.Args()[0]
	port := flag.Args()[1]

	conn, err := net.Dial(protocol, net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal(err)
	}

	signalChan := make(chan os.Signal, 1)
	errorChan := make(chan error, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go send(conn, errorChan)
	go recv(conn, errorChan)

	select {
	case <-signalChan:
		conn.Close()
	case <-errorChan:
		log.Fatal(err)
	}

}

// send бесконечно считывает текси из stdin и отправляет в conn
func send(conn net.Conn, errorChan chan<- error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			errorChan <- err
		}
		fmt.Fprintf(conn, text + "\n")
	}
}

// recv бесконечно считывает текст из conn и выводит в stdout
func recv(conn net.Conn, errorChan chan<- error) {
	reader := bufio.NewReader(conn)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			errorChan <- err
		}
		fmt.Print(text)
	}
}