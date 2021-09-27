package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/rixagis/wb-level-2/develop/dev05/grep"
)

func loadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func main() {
	after := flag.Int("A", 0, "Print num lines of trailing context after matching lines.")
	before := flag.Int("B", 0, "Print num lines of leading context before matching lines. ")
	context := flag.Int("C", 0, "Print num lines of leading and trailing output context.")
	count := flag.Bool("c", false, "Suppress normal output; instead print a count of matching lines for each input file.")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions in patterns and input data, so that characters that differ only in case match each other")
	invert := flag.Bool("v", false, "Invert the sense of matching, to select non-matching lines.")
	fixed := flag.Bool("F", false, "Interpret patterns as fixed strings, not regular expressions.")
	lineNum := flag.Bool("n", false, "Prefix each line of output with the 1-based line number within its input file.")


	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: gerp [OPTIONS]... pattern [FILE]")
		os.Exit(1)
	}

	pattern := flag.Args()[0]
	filename := flag.Args()[1]

	data, err := loadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if *context != 0 {
		*before = *context
		*after = *context
	}

	params := grep.Parameters{
		After: *after,
		Before: *before,
		Count: *count,
		IgnoreCase: *ignoreCase,
		Invert: *invert,
		Fixed: *fixed,
		LineNum: *lineNum,
	}

	grep.Grep(data, pattern, params, os.Stdout)

}