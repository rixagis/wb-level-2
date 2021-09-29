package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/mitchellh/go-ps"
)

// главный цикл шелла
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		prompt()
		line, _ := reader.ReadString('\n')

		if line == "exit\n" {
			break
		}
		if line == "\n" {
			continue
		}
		processForks(line[:len(line)-1])
	}
}

// prompt выводит "приглашение" в начале строки
func prompt() {
	wd, _ := os.Getwd()
	fmt.Print("myshell:", wd, "$ ")
}

// processForks обрабатывает вызовы fork
func processForks(line string) {
	forks := strings.Split(line, "&")
	// форкаем все, кроме последнего
	for i := 0; i < len(forks)-1; i++ {
		ret, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if err != 0 {
			log.Fatal(err)
		}

		// потомок
		if ret == 0 {
			fmt.Println(os.Getpid())
			processPipes(forks[i])
			fmt.Println("Done")
			os.Exit(0)
		}
	}
	// здесь выполняется только родитель
	processPipes(forks[len(forks)-1])
}

// processPipes обрабатывает операторы pipe
func processPipes(line string) {
	commands := strings.Split(line, "|")
	waits := []func() error{}
	out, wait := processCommand(commands[0], os.Stdin)
	waits = append(waits, wait)

	for i := 1; i < len(commands)-1; i++ {
		nextout, wait := processCommand(commands[i], out)
		waits = append(waits, wait)
		out = nextout
	}

	nextout, wait := processCommand(commands[len(commands)-1], out)
	waits = append(waits, wait)
	if nextout != nil {
		_, err := io.Copy(os.Stdout, nextout)
		if err != nil && !errors.Is(err, fs.ErrClosed) {
			log.Fatal(err)
		}
	}
	for _, wait := range waits {
		if wait != nil {
			wait()
		}
	}
}

// processCommand обрабатывает вызов одной команды (команда + аргументы),
// возвращает ридер для чтения результата и функцию ожидания завершения команды
func processCommand(line string, in io.Reader) (io.Reader, func() error) {
	args := strings.Fields(line)
	if len(args) == 0 {
		return nil, nil
	}
	switch args[0] {
	case "pwd":
		return getwd(), nil
	case "cd":
		err := os.Chdir(args[1])
		if err != nil {
			log.Fatal(err)
		}
		return nil, nil
	case "echo":
		str := strings.Join(args[1:], " ")
		return echo(str), nil
	case "ps":
		return processPS(args[1:]), nil
	case "kill":
		for i := 1; i < len(args); i++ {
			pid, err := strconv.Atoi(args[i])
			if err != nil {
				log.Fatal(err)
			}
			err = syscall.Kill(pid, syscall.SIGINT)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil, nil
	default:
		name := args[0]
		cmdargs := []string{}
		if len(args) > 1 {
			cmdargs = args[1:]
		}
		return command(name, in, cmdargs...)
	}
}

// getwd эмулирует команду pwd, возвращает ридер для чтения вывода
func getwd() io.Reader {
	buffer := bytes.Buffer{}
	res, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	buffer.WriteString(res + "\n")
	return &buffer
}

// echo эмулирует одноименную команду, возвращает ридер для чтения вывода
func echo(s string) io.Reader {
	buffer := bytes.Buffer{}
	buffer.WriteString(s + "\n")
	return &buffer
}

// processPS эмулирует команду ps, возвращает ридер для чтения вывода
func processPS(args []string) io.Reader {
	var A bool
	for _, arg := range args {
		switch arg {
		case "-A":
			A = true
		}
	}
	output := bytes.Buffer{}
	output.WriteString("PID\tCMD\n")

	var procs []ps.Process
	procs, err := ps.Processes()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range procs {
		if os.Getppid() == p.Pid() || A {
			output.WriteString(fmt.Sprintf("%d\t%s\n", p.Pid(), p.Executable()))
		}
	}
	return &output
}

// command обрабатывает команду, которой нет в списке реализованных собственноручно,
// возвращает ридер для чтения результата и функцию ожидания завершения команды
func command(name string, in io.Reader, args ...string) (io.Reader, func() error) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = in
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("1", err)
	}

	go func() {
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		/*if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}*/
	}()
	return stdout, cmd.Wait
}
