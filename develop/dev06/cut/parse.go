// функции, парсящие параметр -f
package cut

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var ErrInvalidFieldRange = errors.New("invalid field range")

// ParseIntervals парсит строку с указанием индексов, подаваемую в качестве параметра -f.
// Возвращает слайс точечных индексов полей indices, а также параметры until (-N) и from (N-).
// Если таких параметров не указано, until и from = -1.
// В случае неверного формата вовзвращается ошибка ErrInvalidFieldRange в err
// Слайс indices гарантирует условия единственности и упорядоченности элементов, а также
// отсутствие пересечений с интервалами, указанными в until или from.
// Параметры until и from гарантированно не пересекаются.
func ParseIntervals(input string) (indices []int, until, from int, err error) {
	var (
		intervals = strings.Split(input, ",")
		indexIntervals []string
		fromIntervals []string
		untilIntervals []string
	)

	// разделение по типам интервалов
	for _, interval := range intervals {
		if strings.HasPrefix(interval, "-") {
			untilIntervals = append(untilIntervals, interval)
		} else if strings.HasSuffix(interval, "-") {
			fromIntervals = append(fromIntervals, interval)
		} else {
			indexIntervals = append(indexIntervals, interval)
		}
	}

	// обработка интервалов с индексами (в т.ч. и точечные)
	unnormalizedIndices, err := parseIndexIntervals(indexIntervals)
	if err != nil {
		return nil, -1, -1, err
	}
	// обработка from интервалов
	from, err = parseFromIntervals(fromIntervals)
	if err != nil {
		return nil, -1, -1, err
	}
	// обработка until интервалов
	until, err = parseUntilIntervals(untilIntervals)
	if err != nil {
		return nil, -1, -1, err
	}

	if from > -1 && until > -1 && from <= until {		// покрывают все поля
		return nil, -1, 0, nil
	}

	// нормализация (удаление индексов, заслоненных from или until интервалами)
	for _, index := range unnormalizedIndices {
		if (until == -1 || index > until) && (from == -1 || index < from) {
			indices = append(indices, index)
		}
	}

	return
}

// parseIndexIntervals опереводит интервалы вида N и N-M в слайс единичных индексов.
// Гарантирует единственность и упорядоченность элементов на выходе.
func parseIndexIntervals(intervals []string) ([]int, error) {
	var initial = make(map[int]bool)	// map для соблюдения уникальности номеров полей
	for _, interval := range intervals {
		if strings.Contains(interval, "-") {	// если это не число, а интервал типа m-n
			intervalIndices, err := parseIndexRange(interval)
			if err != nil {
				return nil, ErrInvalidFieldRange
			}
			for _, index := range intervalIndices {
				initial[index] = true
			}
		} else {	// просто число
			n, err := strconv.Atoi(interval)
			if err != nil || n < 1 {
				return nil, ErrInvalidFieldRange
			}
			initial[n] = true
		}
	}

	var indices []int
	for n := range initial {
		indices = append(indices, n)
	}
	sort.Ints(indices)
	return indices, nil
}

// parseIndexRange переводит интервал вида N-M в список единичных индексов
func parseIndexRange(indexRange string) (indices []int, err error) {
	var ends = strings.Split(indexRange, "-")
	if len(ends) != 2 {
		return nil, ErrInvalidFieldRange
	}
	start, err := strconv.Atoi(ends[0])
	if err != nil || start < 1{
		return nil, ErrInvalidFieldRange
	}
	end, err := strconv.Atoi(ends[1])
	if err != nil || end < 1{
		return nil, ErrInvalidFieldRange
	}
	if end < start {
		return nil, ErrInvalidFieldRange
	}

	for i := start; i <= end; i++ {
		indices = append(indices, i)
	}
	return indices, nil
}

// parseIndexRange переводит интервалы вида N- в единственный индекс, покрывающий все эти интервалы
func parseFromIntervals(intervals []string) (int, error) {
	if len(intervals) == 0 {
		return -1, nil
	}
	
	var lowestFrom int
	// первый элемент обрабатывается отдельно, потому что его надо записать в lowestFrom (наименьший)
	_, err := fmt.Sscanf(intervals[0], "%d-", &lowestFrom)
	if err != nil || lowestFrom < 1 {
		return -1, ErrInvalidFieldRange
	}

	for _, interval := range intervals {
		var num int
		_, err := fmt.Sscanf(interval, "%d-", &num)
		if err != nil || num < 1 {
			return -1, ErrInvalidFieldRange
		}
		if num < lowestFrom {
			lowestFrom = num
		}
	}

	return lowestFrom, nil
}

// parseUntilIntervals переводит интервалы вида -N в единственный индекс, покрывающий все эти интервалы
func parseUntilIntervals(intervals []string) (int, error) {
	if len(intervals) == 0 {
		return -1, nil
	}

	var highestUntil = -1
	for _, interval := range intervals {
		var num int
		interval = interval + "END"		// этот трюк нужен, чтобы Sscanf выдавал ошибку, если за числом идет что-то еще
		_, err := fmt.Sscanf(interval, "-%dEND", &num)
		if err != nil || num < 1 {
			return -1, ErrInvalidFieldRange
		}
		if num > highestUntil {
			highestUntil = num
		}
	}

	return highestUntil, nil
}