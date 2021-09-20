package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}

// эта функция возвращает тип *customError, а не error
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()			// В переменную интерфейса error записывается значение nil с динамическим типом *customError
    if err != nil {			// Как и в предыдущем задании, интерфейс равен nil только когда и значение, и тип равны nil,
							// но тут динамический тип равен *customError, поэтому err != nil
        println("error")	// Программа выведет error
        return
    }
    println("ok")
}
