package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run запускает клиент и конкурентные чтение и запись из stdin в соединение и из соединения в stdout.
func Run(timeout time.Duration, host, port string) error {
	
	client := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		return err
	}
	log.Printf("connected to %s\n", client.address)

	defer func(){
		client.Close()
	}()

	signalChan := make(chan os.Signal, 1)
	errorChan := make(chan error, 1)
	doneChan := make(chan struct{})
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go ReceiveWorker(client, errorChan, doneChan)
	go SendWorker(client, errorChan, doneChan)

	defer close(doneChan)

	select {
	case <- signalChan:
		return nil
	case err := <- errorChan:
		return err
	}
}

// SendWorker - горутина, выполняющая операцию Send клиента. При получении ошибки посылает ее в errorChannel и завершает работу.
// Завершает работу при закрытии done
func SendWorker(client *TelnetClient, errorChannel chan<- error, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			err := client.Send()
			if err != nil {
				errorChannel <- err
				return
			}
		}
	}
}

// ReceiveWorker - горутина, выполняющая операцию Receive клиента. При получении ошибки посылает ее в errorChannel и завершает работу.
// Завершает работу при закрытии done
func ReceiveWorker(client *TelnetClient, errorChannel chan<- error, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			err := client.Receive()
			if err != nil {
				errorChannel <- err
				return
			}
		}
	}
}