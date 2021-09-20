package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10 * time.Second, "Таймаут для подключения в секундах")
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalf("Использование: go-telnet [--timeout=<время>] host port")
	}
	if err := Run(*timeout, flag.Args()[0], flag.Args()[1]); err != nil {
		log.Fatalf(err.Error())
	}
}