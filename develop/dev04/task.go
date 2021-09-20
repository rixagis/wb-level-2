package anagrams

import "strings"

// areAnagrams проверяет, являются ли два слова анаграммами
func areAnagrams(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var letters1 = make(map[rune]int)
	var letters2 = make(map[rune]int)

	for _, c := range word1 {
		letters1[c]++
	}
	for _, c := range word2 {
		letters2[c]++
	}

	if len(letters1) != len(letters2) {
		return false
	}

	for c := range letters1 {
		if letters1[c] != letters2[c] {
			return false
		}
	}
	
	return true
}

// MakeAnagramMap создает словарь групп анаграмм, в котором ключем является первое встреченное слово,
// являющееся анаграммой остальных в группе, а значением - слайс слов-анаграмм.
// Все слова в итоговом словаре в нижнем регистре, каждая группа анаграмм содержит не меньше двух членов.
func MakeAnagramMap(input []string) map[string][]string {
	// копируем содержимое массива, чтобы приведение к нижнему регистру не изменило входные данные
	var words = make([]string, len(input))
	copy(words, input)

	// перевод в нижний регистр сразу, так как на выходе по условию должен быть именно он
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	// заполнение словаря
	var anagrams = make(map[string][]string)
	for _, word := range words {
		var added = false
		for key := range anagrams {
			if areAnagrams(key, word) {
				anagrams[key] = append(anagrams[key], word)
				added = true
			}
		}
		if !added {
			anagrams[word] = []string{word}
		}
	}

	// удаление списков из одного элемента
	for key := range anagrams {
		if len(anagrams[key]) == 1 {
			delete(anagrams, key)
		}
	}

	return anagrams
}