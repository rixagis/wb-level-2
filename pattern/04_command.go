// Реализация паттерна "команда" на примере абстрактного (воображаемого) графического редактора
package pattern

import (
	"fmt"
)

// ============================ Вспомогательные типы данных (геометрические фигуры) ================================

// GenericShapeParams - структура с параметрами, которые есть парктически у всех фигур
type GenericShapeParams struct {
	xStart, yStart, xEnd, yEnd int
	color int32
	width int
}

// Line - структура линии
type Line struct {
	GenericShapeParams
}

// Rectangle - структура прямоугольника
type Rectangle struct {
	GenericShapeParams
	fillColor int32
}

// Text - структура текста
type Text struct {
	GenericShapeParams
	text string
}

// ====================================== Команды ==============================================

// Command - интерфейс команды
type Command interface {
	Execute()

	IsSavable() bool	// требуется, чтобы отличить редактирующие команды (рисование) от служебных (undo, redo)
}



// DrawLineCommand - команда рисования линии
type DrawLineCommand struct {
	editor *Editor
	params GenericShapeParams
}

// MakeDrawLineCommand - конструктор DrawLineCommand
func MakeDrawLineCommand(editor *Editor, params GenericShapeParams) DrawLineCommand {
	return DrawLineCommand{editor, params}
}

// Execute выполняет команду
func (c DrawLineCommand) Execute() {
	c.editor.DrawLine(Line{
							c.params,
	})
}

// IsSavable указывает, что эту команду надо хранить в истории
func (c DrawLineCommand) IsSavable() bool {
	return true
}



// DrawRectangleCommand - команда рисования прямоугольника
type DrawRectangleCommand struct {
	editor *Editor
	params GenericShapeParams
	fillColor int32
}

// MakeDrawRectangleCommand - конструктор DrawRectangleCommand
func MakeDrawRectangleCommand(editor *Editor, params GenericShapeParams, fillColor int32) DrawRectangleCommand {
	return DrawRectangleCommand{editor, params, fillColor}
}

// Execute выполняет команду
func (c DrawRectangleCommand) Execute() {
	c.editor.DrawRectangle(Rectangle{
						c.params,
						c.fillColor,
	})
}

// IsSavable указывает, что эту команду надо хранить в истории
func (c DrawRectangleCommand) IsSavable() bool {
	return true
}



// DrawTextCommand - команда рисования текста
type DrawTextCommand struct {
	editor *Editor
	params GenericShapeParams
	text string
}

// MakeDrawTextCommand - конструктор DrawTextCommand
func MakeDrawTextCommand(editor *Editor, params GenericShapeParams, text string) DrawTextCommand {
	return DrawTextCommand{editor, params, text}
}

// Execute выполняет команду
func (c DrawTextCommand) Execute() {
	c.editor.DrawText(Text{
							c.params,
							c.text,
	})
}

// IsSavable указывает, что эту команду надо хранить в истории
func (c DrawTextCommand) IsSavable() bool {
	return true
}



// UndoCommand - команда отмены действия
type UndoCommand struct {
	editor *Editor
}

// MakeUndoCommand - конструктор UndoCommand
func MakeUndoCommand(editor *Editor) UndoCommand {
	return UndoCommand{editor}
}

// Execute выполняет команду
func (c UndoCommand) Execute() {
	c.editor.Undo()
}

// IsSavable указывает, что эту команду НЕ надо хранить в истории
func (c UndoCommand) IsSavable() bool {
	return false
}



// RedoCommand - команда возобновления отмененного действия
type RedoCommand struct {
	editor *Editor
}

// MakeRedoCommand - конструктор RedoCommand
func MakeRedoCommand(editor *Editor) RedoCommand {
	return RedoCommand{editor}
}

// Execute выполняет команду
func (c RedoCommand) Execute() {
	c.editor.Redo()
}

// IsSavable указывает, что эту команду НЕ надо хранить в истории
func (c RedoCommand) IsSavable() bool {
	return false
}



// ====================================== Редактор =====================================================

// Editor - воображаемый графический редактор
type Editor struct {
	commandQueue []Command			// история команд
	lastCommand int					// номер последней активной команды
	currentShapes []interface{}		// "состояние" редактора (слайс нарисованных фигур)
}

// NewEditor - конструктор Editor
func NewEditor() *Editor {
	return &Editor{make([]Command, 0), -1, make([]interface{}, 0)}
}

// DrawLine рисует линию
func (editor *Editor) DrawLine(l Line) {
	editor.currentShapes = append(editor.currentShapes, l)
}

// DrawRectangle рисует прямоугольник
func (editor *Editor) DrawRectangle(r Rectangle) {
	editor.currentShapes = append(editor.currentShapes, r)
}

// DrawText рисует текст
func (editor *Editor) DrawText(t Text) {
	editor.currentShapes = append(editor.currentShapes, t)
}

// ExecuteCommand выполняет (и сохраняет) команду
func (editor *Editor) ExecuteCommand(c Command) {
	if c.IsSavable() {
		if editor.lastCommand < len(editor.commandQueue) - 1 {
			editor.commandQueue = editor.commandQueue[0:editor.lastCommand + 1]
		}
		editor.commandQueue = append(editor.commandQueue, c)
		editor.lastCommand++
	}
	c.Execute()
}

// Undo отменяет последнюю сохраненную команду
func (editor *Editor) Undo() {
	if editor.lastCommand > -1 {
		editor.lastCommand--
		editor.currentShapes = editor.currentShapes[0 : len(editor.currentShapes) - 1]
	}
}

// Redo возобновляет последнюю отмененную команду
func (editor *Editor) Redo() {
	if editor.lastCommand < len(editor.commandQueue) - 1 {
		editor.lastCommand++
		editor.commandQueue[editor.lastCommand].Execute()
	}
}



// reportState - вспомогательная функция для вывода состояния редактора, сделана для целей демонстрации
func reportState(editor *Editor) {
	fmt.Print("Current state:")
	fmt.Println(editor.currentShapes)
}



// ========================================= main ===================================
// Пример использования
/*func main() {
	var editor = NewEditor()

	reportState(editor)

	fmt.Println("\nAdding a line:")
	var lineCommand = MakeDrawLineCommand(
						editor,
						GenericShapeParams{
							xStart: 1, yStart: 1, xEnd: 15, yEnd: 20,
							color: 0x000000,
							width: 1,
						})
	editor.ExecuteCommand(lineCommand)
	reportState(editor)

	fmt.Println("\nAdding text:")
	var textCommand = MakeDrawTextCommand(
						editor,
						GenericShapeParams{
							xStart: 1, yStart: 1, xEnd: 15, yEnd: 20,
							color: 0x000000,
							width: 1,
						},
						"Hello Words")
	editor.ExecuteCommand(textCommand)
	reportState(editor)

	fmt.Println("\nAdding a rectangle:")
	var rectCommand = MakeDrawRectangleCommand(
						editor,
						GenericShapeParams{
							xStart: 1, yStart: 1, xEnd: 15, yEnd: 20,
							color: 0x000000,
							width: 1,
						},
						0x00FF00)
	editor.ExecuteCommand(rectCommand)
	reportState(editor)

	fmt.Println("\nUndoing:")
	var undoCommand = MakeUndoCommand(editor)
	editor.ExecuteCommand(undoCommand)
	reportState(editor)

	fmt.Println("\nRedoing:")
	var redoCommand = MakeRedoCommand(editor)
	editor.ExecuteCommand(redoCommand)
	reportState(editor)

}*/