package main

import (
	"fmt"
	"sync"
	"time"
	"zzzz/myfile"
)

func init() {
	fmt.Println("init")
}
func foo(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	fmt.Println("foo")
	wg.Done()

}
func bar(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	fmt.Println("bar")
	wg.Done()
}
func main() {
	myfile.FileTest()
}
