package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/rixagis/wb-level-2/develop/dev03/sort"
)

func main() {
	var (
		numerical = flag.Bool("n", false, "Сортировать как числа")
		month = flag.Bool("M", false, "Сортировать как месяцы")
		human = flag.Bool("h", false, "Сортировать как числа с суффиксами СИ")
		reverse = flag.Bool("r", false, "Сортировать в обратном порядке")
		unique = flag.Bool("u", false, "Не выводить повторяющиеся строки")
		ignoreTrailingWhitespace = flag.Bool("b", false, "Игнорировать хвостовые пробелы")
		key = flag.Int("k", -1, "Номер колонки, по которой вести сортировку")
		check = flag.Bool("c", false, "Проверить, отсортированы ли строки")
	)
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var filedata = make(map[string][]string, len(flag.Args()))

	for _, filename := range flag.Args() {
		file, err := os.Open(filename)
		if err != nil {
    		fmt.Fprintf(os.Stderr, "could not read file %q: %s\n", filename, err)
			os.Exit(2)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			filedata[filename] = append(filedata[filename], scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "could not read file %q: %s\n", filename, err)
			os.Exit(2)
		}
	}


	if *check {
		for filename, lines := range filedata {
			var result = sort.Check(lines, *key, *numerical, *month, *human, *reverse, *ignoreTrailingWhitespace)
			if result > 0 {
				fmt.Printf("sort: %s:%d: disorder: %s\n", filename, result, lines[result])
				os.Exit(0)
			}
		}
		fmt.Println("All files are sorted")
	}
	
	var allLines []string
	for _, lines := range filedata {
		allLines = append(allLines, lines...)
	}
	allLines = sort.Sort(allLines, *key, *numerical, *month, *human, *reverse, *unique, *ignoreTrailingWhitespace)
	for _, line := range allLines {
		fmt.Println(line)
	}
}