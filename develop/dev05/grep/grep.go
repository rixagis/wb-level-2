package grep

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// Parameters - структура для хранения параметров выполнения grep
type Parameters struct {
	After      int
	Before     int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

// containsPattern определяет, содержится ли подстрока, подходящая под pattern, в str
func containsPattern(str string, pattern string, ignoreCase bool) bool {
	if ignoreCase {
		pattern = "(?i)" + pattern
	}
	res, _ := regexp.MatchString(pattern, str)
	return res
}

// containsRaw определяет, содержится ли подстрока substr в str
func containsRaw(str string, substr string, ignoreCase bool) bool {
	if ignoreCase {
		return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
	}
	return strings.Contains(str, substr)
}

// findIndices находит список индексов всех строк, в которых встречается подстрока или паттерн target,
// в соответствии с параметрами params
func findIndices(data []string, target string, params Parameters) []int {
	var contains func(string, string, bool) bool
	if params.Fixed {
		contains = containsRaw
	} else {
		contains = containsPattern
	}

	res := []int{}
	for i, line := range data {
		if contains(line, target, params.IgnoreCase) != params.Invert {
			res = append(res, i)
		}
	}

	return res
}

// addRangeIndices добавляет в список индексов те индексы, которые входят в контекст (флаги -A, -B, -C)
func addRangeIndices(indices []int, length int, params Parameters) []int {
	uniques := make(map[int]struct{})
	
	for _, index := range indices {
		start := 0
		if index - params.Before >= 0 {
			start = index - params.Before
		}

		end := length
		if index + params.After < length {
			end = index + params.After + 1
		}

		for i:= start; i < end; i++ {
			uniques[i] = struct{}{}
		}
	}
	
	res := make([]int, 0, len(uniques))
	for index := range uniques {
		res = append(res, index)
	}
	sort.Ints(res)

	return res
}

// Grep производит поиск по слайсу строк на предмет подстроки или паттерна target во соответствии с параметрами params.
// Вывод осуществляется в out.
func Grep(data []string, target string, params Parameters, out io.Writer) {
	indices := findIndices(data, target, params)

	if params.Count {
		fmt.Fprintf(out, "%d\n", len(indices))
		return
	}

	indices = addRangeIndices(indices, len(data), params)

	if params.LineNum {
		for _, index := range indices {
			fmt.Fprintf(out, "%d:%s\n", index+1, data[index])
		}
	} else {
		for _, index := range indices {
			fmt.Fprintf(out, "%s\n", data[index])
		}
	}
}