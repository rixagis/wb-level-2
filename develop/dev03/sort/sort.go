package sort

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// sortByWhole сортирует слайс arr с помощью функции сравнения less.
// флаги:
// reversed - сортировать в обратном порядке
// ignoreTrailingWhitespace - игнорировать пробелы ы конце
func sortByWhole(arr []string, less func([]string, int, int, bool) bool, reversed, ignoreTrailingWhitespace bool) {
	sort.SliceStable(arr, func(i, j int) bool {
		if reversed {
			return less(arr, j, i, ignoreTrailingWhitespace)
		}else{
			return less(arr, i, j, ignoreTrailingWhitespace)
		}
	})
}

// sortByField сортирует слайс arr с помощью функции сравнения less по колонке с номером fieldNum
// флаги:
// reversed - сортировать в обратном порядке
// ignoreTrailingWhitespace - игнорировать пробелы ы конце
func sortByField(arr []string, fieldNum int, less func([]string, int, int, bool) bool, reversed bool) {
	sort.SliceStable(arr, func(i, j int) bool {
		if reversed {
			i, j = j, i
		}
		var fields1 = strings.Fields(arr[i])
		var fields2 = strings.Fields(arr[j])

		if len(fields1) < fieldNum {
			if len(fields2) < fieldNum {
				return less(arr, i, j, false)
			} else {
				return true
			}
		} else {
			if len(fields2) < fieldNum {
				return false
			} else {
				return less([]string{fields1[fieldNum - 1], fields2[fieldNum - 1]}, 0, 1, false)
			}
		}
	})
}

// lessLexicographical - функция сравнения для лексикографического порядка
func lessLexicographical(arr []string, i, j int, ignoreTrailingWhitespace bool) bool {
	var a = arr[i]
	var b = arr[j]
	if ignoreTrailingWhitespace {
		a = strings.TrimRightFunc(a, unicode.IsSpace)
		b = strings.TrimRightFunc(b, unicode.IsSpace)
	}
	return a < b
}


// getNumberPart возвращает число, стоящее в начале строки, если такое есть, иначе - 0
func getNumberPart(s string) float64 {
	var l = len(s)
	if len(s) == 0 {
		return 0
	}
	var runes = []rune(s)
	var start = 0
	var numberBuilder = strings.Builder{}
	if runes[0] == '-' {
		if l == 1 {
			return 0
		}
		start = 1
		numberBuilder.WriteString("-")
	}

	var passedDot = false
	for i := start; i < l; i++ {
		if runes[i] == '.' {
			if passedDot {
				break
			} else {
				passedDot = true
				numberBuilder.WriteRune('.')
			}
		} else if runes[i] >= '0' && runes[i] <= '9' {
			numberBuilder.WriteRune(runes[i])
		} else {
			break
		}
	}

	var res, _ = strconv.ParseFloat(numberBuilder.String(), 64)
	return res
}

// lessNumerical - функция сравнения для числел
func lessNumerical(arr []string, i, j int, ignoreTrailingWhitespace bool) bool {
	var a = getNumberPart(arr[i])
	var b = getNumberPart(arr[j])
	if a == b {
		return lessLexicographical(arr, i, j, ignoreTrailingWhitespace)
	} else {
		return a < b
	}
}

// проверка наличия строки target в слайсе arr
func in(target string, arr []string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}

// фильтрует arr и возвращает только уникальные строки
func getUniques(arr []string) []string {
	var l = len(arr)
	var uniques = make([]string, 0, l)
	for _, s := range arr {
		if !in(s, uniques){
			uniques = append(uniques, s)
		}
	}
	return uniques
}

// переводит название месяца в его порядковый номер
func monthToInt(month string) int {
	month = strings.ToLower(month)
	switch month {
	case "jan", "january":
		return 1
	case "feb", "february":
		return 2
	case "mar", "march":
		return 3
	case "apr", "april":
		return 4
	case "may":
		return 5
	case "jun", "june":
		return 6
	case "jul", "july":
		return 7
	case "aug", "august":
		return 8
	case "sep", "september":
		return 9
	case "oct", "october":
		return 10
	case "nov", "november":
		return 11
	case "dec", "december":
		return 12
	default:
		return -1
	}
}

// функция сравнения для месяцев
func lessMonth(arr []string, i, j int, ignoreTrainigWhitespace bool) bool {
	var (
		a = arr[i]
		b = arr[j]
	)

	if ignoreTrainigWhitespace {
		a = strings.TrimRightFunc(a, unicode.IsSpace)
		b = strings.TrimRightFunc(b, unicode.IsSpace)
	}
	return monthToInt(a) < monthToInt(b)
}

// получает суффикс, идущий за числом, стоящим в начале строки
func getSuffix(s string) string {
	var start = 0
	var runes = []rune(s)
	var l = len(runes)
	if runes[0] == '-' {
		start = 1
	}

	var passedDot = false
	var index = start
	for index < l {
		if runes[index] == '.' {
			if passedDot {
				break
			} else {
				passedDot = true
			}
		} else if runes[index] < '0' || runes[index] > '9' {
			break
		}
		index++
	}
	var result = runes[index:]
	return string(result)
}

// функция сравнения для суффиксов (-h)
func lessSuffix(arr []string, i, j int, ignoreTrailingWhitespace bool) bool {
	var (
		a = arr[i]
		b = arr[j]
	)
	
	if ignoreTrailingWhitespace {
		a = strings.TrimRightFunc(a, unicode.IsSpace)
		b = strings.TrimRightFunc(b, unicode.IsSpace)
	}

	var (
		number1 = getNumberPart(a)
		number2 = getNumberPart(b)
		suffix1 = getSuffix(a)
		suffix2 = getSuffix(b)
		suffixes = "kKMGTPEZY"
		suffix1Value = strings.Index(suffixes, suffix1)
		suffix2Value = strings.Index(suffixes, suffix2)
	)



	if suffix1 == "" {suffix1Value = -1}
	if suffix2 == "" {suffix2Value = -1}

	if suffix1Value == suffix2Value {
		if number1 == number2 {
			return a < b
		} else {
			return number1 < number2
		}
	} else {
		return suffix1Value < suffix2Value
	}
}

// Sort сортирует слайс arr согласно параметрам.
// Параметры:
//  key - номер колонки, по которой надо сортировать, при -1 строки сравниваются целиком
//  numerical - сортировать как числа
//  month - сортировать как месяцы
//  human - сортировать как числа с суффиксами СИ
//  reversed - сортировать в обратном порядке
//  unique - вернуть в результате только уникальные строки
//  ignoreTrailSpaces - игнорировать хвостовые пробелы
func Sort(arr []string, key int, numerical, month, human, reversed, unique, ignoreTrailSpaces bool) []string {
	var lessFunc func ([]string, int, int, bool) bool
	if numerical {
		lessFunc = lessNumerical
	} else if month {
		lessFunc = lessMonth
	} else if human {
		lessFunc = lessSuffix
	} else {
		lessFunc = lessLexicographical
	}

	var result []string

	if unique {
		result = getUniques(arr)
	} else {
		result = append(result, arr...)
	}

	if key > 0 {
		sortByField(result, key, lessFunc, reversed)
	} else {
		sortByWhole(result, lessFunc, reversed, ignoreTrailSpaces)
	}

	return result
}

// Check проверяет, отсортирована ли слайс arr согласно параметрам.
// Возвращает номер первой строки, идущей не по порядку, или -1, если строка отсортирована.
// Параметры:
//  key - номер колонки, по которой надо сортировать, при -1 строки сравниваются целиком
//  numerical - сортировать как числа
//  month - сортировать как месяцы
//  human - сортировать как числа с суффиксами СИ
//  reversed - сортировать в обратном порядке
//  ignoreTrailSpaces - игнорировать хвостовые пробелы
func Check(arr []string, key int, numerical, month, human, reversed, ignoreTrailSpaces bool) int {
	var lessFunc func ([]string, int, int, bool) bool
	if numerical {
		lessFunc = lessNumerical
	} else if month {
		lessFunc = lessMonth
	} else if human {
		lessFunc = lessSuffix
	} else {
		lessFunc = lessLexicographical
	}

	var l = len(arr)


	for i := 1; i < l; i++ {
		var a = arr[i - 1]
		var b = arr[i]

		if key > 0 {
			var fields1 = strings.Fields(a)
			var fields2 = strings.Fields(b)

			if len(fields1) >= key {
				a = fields1[key - 1]
			}
			if len(fields2) >= key {
				b = fields2[key - 1]
			}
		}

		if reversed {
			a, b = b, a
		}

		if lessFunc([]string{b, a}, 0, 1, ignoreTrailSpaces) {
			return i+1
		}
	}

	return -1
}
