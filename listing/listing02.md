package main
 
import (
    "fmt"
)
 
 
// функция с именованным результатом, функция в defer может изменять именованный результат
func test() (x int) {
    defer func() {		// выполняется после return, увеличивает результат на 1
        x++
    }()
    x = 1
    return
}
 
 // обычная функция без именованного результата
func anotherTest() int {
    var x int
	// функция в defer выполнится после return, но результат не изменится, так как он не именованный
    defer func() {
        x++
    }()
    x = 1
    return x
}
 
 
func main() {
    fmt.Println(test())			// выводит 2 (именованный результат увеличивается)
    fmt.Println(anotherTest())	// выводит 1 (безымянный результат не меняется)
}
