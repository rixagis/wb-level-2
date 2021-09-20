package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {					// error - интерфейс
    var err *os.PathError = nil		// err имеет тип *os.PathError и значение nil
    return err
}
 
func main() {
    err := Foo()					// err - интерфейс error, содержащий динамический тип *os.PathError и значение nil
    fmt.Println(err)				// выводит <nil>, так как имеет именно такое значение
    fmt.Println(err == nil)			// выводит false, так как сравнение интерфейса с nil возвзращает истину только если
									// в переменной интерйеса и значение, и динамический тип nil
}