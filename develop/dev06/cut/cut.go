// функции, реализующие основную логику программы (разделение строки по колонкам)
package cut

import (
	"strings"
)

// CutLine разделяет строку line по разделителю и возвращает строку,
// состоящую из полей, указанных в индексах, соедененных тем же разделителем.
// Результат ok показывает, что эту строку нужно отобразить, даже если она пустая (не игнорируется).
// Параметры:
//  line - строка, которую нужно разделить
//  indices - слайс индексов полей, который надо вывести 
//  until - указание полей от начала строки до этого
//  from - указание полей от этого, до конца строки
//  delimeter - разделитель полей
//  strict - игнорировать строки без разделителей
// Слайс indices должен удовлетворять условиям уникальности и упорядоченнсти элементов,
//  а также индексы не должны попадать в интервалы, покрываемые from и until.
// Параметры from и until не должны пересекаться.
// Если параметры from или until не требуются, их нужно указать -1
func CutLine(line string, indices []int, until int, from int, delimeter string, strict bool) (result string, ok bool) {
	if !strings.Contains(line, delimeter) {
		if strict {
			return "", false
		} else {
			return line, true
		}
	}

	var fields = strings.Split(line, delimeter)
	var builder = strings.Builder{}

	// Обработка until (эти поля идут в начале)
	if until > -1 {
		builder.WriteString(cutUntil(fields, until, delimeter))
	}

	// обычные индексы (идут между until и from)
	var body = cutIndices(fields, indices, delimeter)

	if builder.String() != "" && body != "" {
		builder.WriteString(delimeter)
	}

	builder.WriteString(body)

	// обработка from (эти поля идут в конце)
	if from > -1 {	// индекс from указан
		var tail = cutFrom(fields, from, delimeter)
		if tail != "" {
			if builder.String() != "" {
				builder.WriteString(delimeter)
			}
			builder.WriteString(tail)
		}
	}

	return builder.String(), true
}

// cutUntil обрабатывает индекс until, возвращая промежуточный результат в виде строки
func cutUntil(fields []string, until int, delimeter string) string {
	var limit int
	if len(fields) - 1 < until {
		limit = len(fields)
	} else {
		limit = until + 1
	}

	return strings.Join(fields[:limit], delimeter)
}

// cutFrom обрабатывает индекс from, возвращая промежуточный результат в виде строки
func cutFrom(fields []string, from int, delimeter string) string {
	var length = len(fields)
	if from >= length || from < 0 {
		return ""
	}

	return strings.Join(fields[from:], delimeter)
}

// cutIndices обрабатывает индексы, не являющиеся from или until, возвращая промежуточный результат в виде строки
func cutIndices(fields []string, indices []int, delimeter string) string {
	var (
		builder = strings.Builder{}
		fieldsLength = len(fields)
		indicesLength = len(indices)
	)

	if indicesLength > 0 && fieldsLength > indices[0] {
		builder.WriteString(fields[indices[0]])
	}

	for i := 1; i < len(indices); i++ {
		if fieldsLength > indices[i] {
			builder.WriteString(delimeter)
			builder.WriteString(fields[indices[i]])
		}
	}

	return builder.String()
}