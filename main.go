package main

import (
	"fmt"
	"time"
	"zzzz/myrpc"
)
func doSomething() {
	go func() {
		cnt := 0
		for {
			cnt ++
			fmt.Println(cnt)
			time.Sleep(time.Second)
		}
	}()
}
func main() {
	go myrpc.RpcServer()
	go myrpc.RpcClient()

	ch := make(chan int)
	ch <- 10
}