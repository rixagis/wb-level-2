package unpacker



// Обозначения состояний автомата
const (
	stateStartID = iota		// начальное стостояние
	stateCharID				// последним был считан символ (не цифра и не \)
	stateNumberID			// последним была считана цифра
	stateBackslashID		// последним был считан \
	stateErrorID			// конечное состояние, указывающее на ошибку формата
	stateFinishID			// конечное успешное состояние
)

type state interface {
	id() int					// id состояния
	processChar(rune) string	// обработка одного входного символа
	setData(rune, int)			// установка дополнительных данных (последние считанные символ и число)
	setContext(*Unpacker)		// связь с родительской структурой
}

// данные, общие для всех состояний
type stateData struct {
	context *Unpacker	// родительская структура
	currentChar rune	// последний считанный символ
	currentNumber int	// последнее считанное число
}

// ========================= START =========================

// начальное состояние
type stateStart struct {
	stateData
}

func (s *stateStart) id() int {
	return stateStartID
}

func (s *stateStart) processChar(c rune) string {
	if _, isDigit := atoi(c); isDigit {						// цифра
		s.context.setState(s.context.errorState, 0, 0)
		return ""
	} else if c == '\\' {									// бэкслеш
		s.context.setState(s.context.backslashState, 0, 0)
		return ""
	} else if c == 0 {										// конец строки
		s.context.setState(s.context.finishState, 0, 0)
		return ""
	} else {
		s.context.setState(s.context.charState, c, 0)		// другой символ
		return ""
	}
}

func (s *stateStart) setData(nextChar rune, nextNumber int) {
}

func (s *stateStart) setContext(c *Unpacker) {
	s.context = c
}



// =================================== CHAR ==============================================

// состояние после прочтения символа (не числа и не \)
type stateChar struct {
	stateData
}

func (s *stateChar) id() int {
	return stateCharID
}

func (s *stateChar) processChar(c rune) (output string) {
	if n, isDigit := atoi(c); isDigit {	// цифра
		s.context.setState(s.context.numberState, s.currentChar, n)
		return ""
	} else if c == '\\' {				// бэкслеш
		s.context.setState(s.context.backslashState, 0, 0)
		return string(s.currentChar)
	} else if c == 0 {					// конец строки
		output = string(s.currentChar)
		s.context.setState(s.context.finishState, 0, 0)
		return output
	} else {							// другой символ
		output = string(s.currentChar)
		s.context.setState(s.context.charState, c, 0)
		return output
	}
}

func (s *stateChar) setData(nextChar rune, nextNumber int) {
	s.currentChar = nextChar
	s.currentNumber = nextNumber
}

func (s *stateChar) setContext(c *Unpacker) {
	s.context = c
}



// =============================== NUMBER ========================================

// состояние после прочтения цифры (идет накопление числа)
type stateNumber struct {
	stateData
}

func (s* stateNumber) id() int {
	return stateNumberID
}

func (s* stateNumber) processChar(c rune) (output string) {
	if n, isDigit := atoi(c); isDigit {						// цифра
		s.currentNumber *= 10
		s.currentNumber += n
		s.context.setState(s.context.numberState, s.currentChar, s.currentNumber)
		return ""
	} else if c == '\\' {									// бэкслеш
		s.context.setState(s.context.backslashState, 0, 0)
		return multiplyRune(s.currentChar, s.currentNumber)
	} else if c == 0 {										// конец строки
		output = multiplyRune(s.currentChar, s.currentNumber)	
		s.context.setState(s.context.finishState, 0, 0)
		return output
	} else {												// другой символ
		output = multiplyRune(s.currentChar, s.currentNumber)
		s.context.setState(s.context.charState, c, 0)
		return output
	}
}

func (s *stateNumber) setData(nextChar rune, nextNumber int) {
	s.currentChar = nextChar
	s.currentNumber = nextNumber
}

func (s *stateNumber) setContext(c *Unpacker) {
	s.context = c
}



// ====================================== BACKSLASH ======================================

// состояние после прочтения \ (идет обработка escape-комбинации)
type stateBackslash struct {
	stateData
}

func (s* stateBackslash) id() int {
	return stateBackslashID
}

func (s* stateBackslash) processChar(c rune) (output string) {
	if _, isDigit := atoi(c); isDigit || c == '\\' {	// цифра или еще один бэкслеш
		s.context.setState(s.context.charState, c, 0)
		return ""
	} else if c == 0 {									// конец строки
		s.context.setState(s.context.errorState, 0, 0)
		return ""
	} else {											// другой символ
		s.context.setState(s.context.errorState, 0, 0)
		return ""
	}
}

func (s *stateBackslash) setData(nextChar rune, nextNumber int) {
	s.currentChar = nextChar
	s.currentNumber = nextNumber
}

func (s *stateBackslash) setContext(c *Unpacker) {
	s.context = c
}

// ==================================== ERROR ==============================================

// ошибочной состояние, сюда автомат заходит при некорректном формате ввода
type stateError struct {
	stateData	
}

func (s *stateError) id() int {
	return stateErrorID
}

func (s *stateError) processChar(c rune) string {
	return ""
}

func (s *stateError) setData(nextChar rune, nextNumber int) {
}

func (s *stateError) setContext(c *Unpacker) {
	s.context = c
}

// ============================================ FINISH ==============================================

// финальное успешное состояние, сюда автомат попадает после успешного прочтения всех строки
type stateFinish struct {
	stateData	
}

func (s *stateFinish) id() int {
	return stateFinishID
}

func (s *stateFinish) processChar(c rune) string {
	s.context.setState(s.context.errorState, 0, 0)
	return ""
}

func (s *stateFinish) setData(nextChar rune, nextNumber int) {
}

func (s *stateFinish) setContext(c *Unpacker) {
	s.context = c
}