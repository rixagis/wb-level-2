package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	
	"github.com/rixagis/wb-level-2/develop/dev06/cut"
)


func main() {
	var (
		fields = flag.String("f", "", "Список номеров колонок или их интервалов (включая открытые)")
		delimiter = flag.String("d", "\t", "Разделитель колонок")
		strict = flag.Bool("s", false, "Игнорировать строки без разделителя")
	)
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Fprintf(os.Stderr, "Неизвестный параметр: %s\n", flag.Args()[0])
		os.Exit(1)
	}

	if *fields == "" {
		fmt.Fprintln(os.Stderr, "Параметр -f обязателен")
		os.Exit(1)
	}

	var fieldIndices, until, from, err = cut.ParseIntervals(*fields)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Неверный формат записи интервалов")
		os.Exit(1)
	}

	for i := range fieldIndices {
		fieldIndices[i]-- // единица вычитается для приведения к системе отсчтета от нуля
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// единицы вычитаются для приведения к системе отсчтета от нуля
		var result, ok = cut.CutLine(scanner.Text(), fieldIndices, until-1, from-1, *delimiter, *strict)
		if ok {
			fmt.Println(result)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}