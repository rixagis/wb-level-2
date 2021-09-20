package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

// TelnetClient - клиент, выполняющий элементарные операции передачи данных по сети
type TelnetClient struct {
	address		string
	timeout		time.Duration
	conn		net.Conn
	inputReader *bufio.Reader
	connReader	*bufio.Reader
	out			io.Writer
}

// NewTelnetClient - конструктор TelnetClient
// Параметры:
//  address - адрес соединения
//  timeout - таймаут соединения
//  in - io.Reader для получения данных, которые будут отправляться в соединение
//  out - io.Writer для посылки данных, получаемых из соединения
func NewTelnetClient(address string, timeout time.Duration, in io.Reader, out io.Writer) *TelnetClient {
	return &TelnetClient{
		address: address,
		timeout: timeout,
		inputReader: bufio.NewReader(in),
		out: out,
	}
}

// Connect производит соединение по указанному в конструкторе адресу
func (t *TelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return err
	}
	t.conn = conn
	t.connReader = bufio.NewReader(t.conn)
	return nil
}

// Receive выполняет чтение из соединения в out
func (t *TelnetClient) Receive() error {
	receved, err := t.connReader.ReadByte()
	if err != nil {
		return err
	}

	if _, err := fmt.Fprint(t.out, string(receved)); err != nil {
		return err
	}
	return nil
}

// Send выполняет чтение из in в соединение
func (t *TelnetClient) Send() error {
	line, err := t.inputReader.ReadString('\n')
	fmt.Println("READ:", line)
	if err != nil {
		if err == io.EOF {	// Соединение закрывается пользователем
			return nil
		}
		return err
	}

	_, err = t.conn.Write([]byte(line))
	return err
}

// Close закрывает соединение
func (t *TelnetClient) Close() error {
	err := t.conn.Close()
	return err
}