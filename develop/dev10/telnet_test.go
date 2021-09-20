package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"sync"
	"testing"
	"time"
)

func TestTelnetClient(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		l, err := net.Listen("tcp", "127.0.0.1:")

		if err != nil {
			t.Errorf("testing basic, expected no error on listen, got: %s", err)
		}
		
		defer func() {
			err := l.Close()
			if err != nil {
				t.Errorf("testing basic, expected no error on listener close, got: %s", err)
			}
		}()

		var wg sync.WaitGroup
		wg.Add(2)

		// клиент
		go func() {
			defer wg.Done()

			in := &bytes.Buffer{}
			out := &bytes.Buffer{}

			timeout := time.Duration(10 * time.Second)
			

			client := NewTelnetClient(l.Addr().String(), timeout, ioutil.NopCloser(in), out)
			
			err = client.Connect()
			if err != nil {
				t.Errorf("testing basic, expected no error, got: %s", err)
			}
			
			defer func() {
				err := client.Close()
				if err != nil {
					t.Errorf("testing basic, expected no error on client close, got: %s", err)
				}
			}()

			in.WriteString("hello\n")
			err = client.Send()
			
			if err != nil {
				t.Errorf("testing basic, expected no error on clint send, got: %s", err)
			}

			for i := 0; i < 6; i++{
				err = client.Receive()
				if err != nil {
					t.Errorf("testing basic, expected no error on client receive, got: %s", err)
				}
			}

			var expected = "world\n"
			var result = out.String() 
			if result != "world\n" {
				t.Errorf("testing basic, expected: %s, got: %s", expected, result)
			}
		}()

		// сервер
		go func() {
			defer wg.Done()

			conn, err := l.Accept()
			if err != nil {
				t.Errorf("testing basic, expected no error on accept, got: %s", err)
			}
			if conn == nil {
				t.Errorf("testing basic, expected conn to be not nil, got nil")
			}
			defer func() {
				err := conn.Close()
				if err != nil {
					t.Errorf("testing basic, expected no error on conn close, got: %s", err)
				}
			}()

			request := make([]byte, 1024)
			n, err := conn.Read(request)
			if err != nil {
				t.Errorf("testing basic, expected no error on conn.Read, got: %s", err)
			}
			var expected = "hello\n"
			var result = string(request)[:n]
			if expected != result {
				t.Errorf("testing basic, expected: %s, got: %s", expected, result)
			}

			n, err = conn.Write([]byte("world\n"))
			if err != nil {
				t.Errorf("testing basic, expected no error, got: %s", err)
			}
			if n != 6 {
				t.Errorf("testing basic, expected: %d, got: %d", 6, n)
			}
		}()

		wg.Wait()
	})

	t.Run("should return error when connect to wrong host", func(t *testing.T) {
		wrongAddress := "localhost:909090"
		in := &bytes.Buffer{}
		out := &bytes.Buffer{}

		timeout, err := time.ParseDuration("10s")
			
		if err != nil {
			t.Errorf("testing basic, expected no error on parse duration, got: %s", err)
		}

		client := NewTelnetClient(wrongAddress, timeout, ioutil.NopCloser(in), out)

		err = client.Connect()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("workers should work correctly", func(t *testing.T) {
		l, err := net.Listen("tcp", "127.0.0.1:")

		if err != nil {
			t.Errorf("testing workers, expected no error on listen, got: %s", err)
		}
		
		defer func() {
			err := l.Close()
			if err != nil {
				t.Errorf("testing workers, expected no error on listener close, got: %s", err)
			}
		}()

		var wg sync.WaitGroup
		wg.Add(2)

		// сервер
		go func() {
			defer wg.Done()

			conn, err := l.Accept()
			if err != nil {
				t.Errorf("testing workers, expected no error on accept, got: %s", err)
			}
			if conn == nil {
				t.Errorf("testing workers, expected conn to be not nil, got nil")
			}
			defer func() {
				err := conn.Close()
				if err != nil {
					t.Errorf("testing workers, expected no error on conn close, got: %s", err)
				}
			}()

			request := make([]byte, 1024)
			n, err := conn.Read(request)
			if err != nil {
				t.Errorf("testing workers, expected no error on conn.Read, got: %s", err)
			}
			var expected = "hello\n"
			var result = string(request)[:n]
			if expected != result {
				t.Errorf("testing workers, expected: %s, got: %s", expected, result)
			}

			n, err = conn.Write([]byte("world\n"))
			if err != nil {
				t.Errorf("testing workers, expected no error, got: %s", err)
			}
			if n != 6 {
				t.Errorf("testing workers, expected: %d, got: %d", 6, n)
			}
		}()

		// клиент
		go func() {
			defer wg.Done()

			realStdin := os.Stdin
			file, _ := ioutil.TempFile(os.TempDir(), "stdin")
			defer os.Remove(file.Name())

			file.WriteString("hello\n")
			file.Seek(0, 0)
			os.Stdin = file


			err = Run(time.Duration(10 * time.Second), "localhost", fmt.Sprint(l.Addr().(*net.TCPAddr).Port))

			if err != nil && err != io.EOF {
				t.Errorf("testing workers, expected no error, got: %s", err)
			}

			os.Stdin = realStdin
		}()

		wg.Wait()
	})
}