package unpacker

// atoi переводит символ цифры в int, если это возможно, или возвращает -1, если невозможно.
// второй результат (bool) указывает, является ли символ цифрой.
func atoi(c rune) (int, bool) {
	if c < '0' || c > '9' {
		return -1, false
	}
	return int(c - '0'), true
}

// multiplyRune возвращает строку с символом c, повторенным n раз
func multiplyRune(c rune, n int) string {
	if n < 0 {
		return ""
	}
	var runes = make([]rune, n)
	for i := range runes {
		runes[i] = c
	}
	return string(runes)
}