package main

import (
	"fmt"
	"sync"
	"time"
)

type A struct {
	a []string
}

func doSomething(wg *sync.WaitGroup) {
	x := 0
	for i := 0; i < 10; i++ {
		x++
		fmt.Printf("work %v\n", x)
		time.Sleep(time.Second)
	}
	wg.Done()
}
func f(ch chan int) {
	x := <- ch
	fmt.Println(x)
}
func main() {
	ch := make(chan int, 10)
	ch <- 10
	fmt.Println(<- ch)
}
