package main
 
import (
    "fmt"
    "math/rand"
    "time"
)

// эта функция создает канал и горутину, которая записывает в канал все аргументы, после чего закрывает канал
func asChan(vs ...int) <-chan int {
   c := make(chan int)
 
   go func() {
       for _, v := range vs {
           c <- v
           time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      }
 
      close(c)
  }()
  return c
}


func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {					// этот селект никак не обрабатывает закрытие каналов,
               case v := <-a:		// поэтому когда какой-либо из каналов закроется - 
                   c <- v			// он будет бесконечно генерировать 0
              case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}
 
func main() {
 
   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)		// выводит числа 1-8 в случайном порядке, после чего бесконечно выводит 0,
	   						// потому что merge не обрабатывает закрытие каналов
   }
}
