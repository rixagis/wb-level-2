package unpacker

import (
	"errors"
	"strings"
)

// Unpacker - автомат, обрабатывающий строку и распаковывающий ее согласно алгоритму:
// "a4bc2d5e" => "aaaabccddddde"
// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "" => ""
type Unpacker struct {
	currentState state		// текущее состояние

	// все состояния, хранятся в структуре, чтобы не приходилось каждый раз создавать новое
	startState     state	// начальное состояние
	charState      state	// после прочтения символа (не цифиры и не \)
	numberState    state	// после прочтения числа
	backslashState state	// после прочтения \
	errorState     state	// ошибка
	finishState    state	// конечное успешное состояние
}

// NewUnpacker - конструктор Unpacker
func NewUnpacker() *Unpacker {
	var p = &Unpacker{}

	p.startState = &stateStart{}
	p.startState.setContext(p)

	p.charState = &stateChar{}
	p.charState.setContext(p)

	p.numberState = &stateNumber{}
	p.numberState.setContext(p)

	p.backslashState = &stateBackslash{}
	p.backslashState.setContext(p)

	p.errorState = &stateError{}
	p.errorState.setContext(p)

	p.finishState = &stateFinish{}
	p.finishState.setContext(p)

	p.setState(p.startState, 0, 0)

	return p
}

// Переход к новому состоянию с указанынми данными
func (p *Unpacker) setState(nextState state, nextChar rune, nextNumber int) {
	p.currentState = nextState
	p.currentState.setData(nextChar, nextNumber)
}

// Обработка одного входного символа, возвращает вывод автомата
func (p *Unpacker) processChar(c rune) string {
	return p.currentState.processChar(c)
}

// Возвращает ID текущего состояния
func (p *Unpacker) currentStateID() int {
	return p.currentState.id()
}

// Unpack распаковывает строку:
// "a4bc2d5e" => "aaaabccddddde"
// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "" => ""
func (p *Unpacker) Unpack(input string) (string, error) {
	var builder = strings.Builder{}

	for _, c := range input {
		builder.WriteString(p.processChar(c))
		if p.currentState == p.errorState {
			return "", errors.New("incorrect format")
		}
	}

	// обработка конца строки
	builder.WriteString(p.processChar(0))
	if p.currentState == p.errorState {
		return "", errors.New("incorrect format")
	}

	return builder.String(), nil
}