package main

import (
	"fmt"
	"time"
	"zzzz/mytime"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func Period(ch chan<- int) {
	for {
		fmt.Println("work")
		time.Sleep(3 * time.Second)
	}
}
func main() {
	mytime.TestSelect()
}
