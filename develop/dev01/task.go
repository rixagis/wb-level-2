package main

import (
	"fmt"
	"time"
	"os"

	"github.com/beevik/ntp"
)


// getCurrentTime оборачивает вызов библиотечной функции, возвращает текущее время и ошибку
func getCurrentTime() (time.Time, error) {
	return ntp.Time("0.beevik-ntp.pool.ntp.org")
}

func main() {
	var time, err = getCurrentTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(time.Format("02-01-2006 15:04:05 MST"))	// кастомный формат для более красивого вывода
}