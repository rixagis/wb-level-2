// Реализация паттерна "фабричный метод" на примере абстрактного фреймворка графического интерфейса
package pattern

import (
)

// Panel - родительская структора элементов интерфейса
type Panel struct {
	x, y int
	w, h int
}

// Button - кнопка
type Button struct {
	Panel
	rounded bool
	shadowed bool
}

// TextBox - текстовое поле
type TextBox struct {
	Panel
	rounded bool
	blinkingCursor bool
}

// GUIFactory - интерфейс фабрик
type GUIFactory interface {
	NewButton(x, y, w, h int) *Button
	NewTextBox(x, y, w, h int) *TextBox
}

// FancyGUIFactory - фабрика "красивого" интерфейса
type FancyGUIFactory struct {
}

// NewButton создает "красивую" кнопку
func (f *FancyGUIFactory) NewButton(x, y, w, h int) *Button {
	return &Button{
		Panel: Panel{x, y, w, h},
		rounded: true,
		shadowed: true,
	}
}

// NewTextBox создает "красивый" текстбокс
func (f *FancyGUIFactory) NewTextBox(x, y, w, h int) *TextBox {
	return &TextBox{
		Panel: Panel{x, y, w, h},
		rounded: true,
		blinkingCursor: true,
	}
}




// MinimalisticGUIFactory - фабрика минималистичного интерфейса
type MinimalisticGUIFactory struct {
}

// NewButton создает минималистичную кнопку
func (f *MinimalisticGUIFactory) NewButton(x, y, w, h int) *Button {
	return &Button{
		Panel: Panel{x, y, w, h},
		rounded: false,
		shadowed: false,
	}
}

// NewTextBox создает минималистичный текстбокс
func (f *MinimalisticGUIFactory) NewTextBox(x, y, w, h int) *TextBox {
	return &TextBox{
		Panel: Panel{x, y, w, h},
		rounded: false,
		blinkingCursor: false,
	}
}

// Пример использования
/*func main() {
	var (
		fancyFactory = &FancyGUIFactory{}
		minimalisticFactory = &MinimalisticGUIFactory{}

		fancyTextbox = fancyFactory.NewTextBox(0, 0, 200, 50)
		fancyButton = fancyFactory.NewButton(0, 100, 200, 50)

		minTextbox = minimalisticFactory.NewTextBox(0, 0, 200, 50)
		minButton = minimalisticFactory.NewButton(0, 0, 200, 50)
	)

	fmt.Println("Fancy textbox:", fancyTextbox)
	fmt.Println("Fancy button:", fancyButton)

	fmt.Println("Minimalistic textbox:", minTextbox)
	fmt.Println("Minimalistic button:", minButton)
}*/