package main

import (
	"fmt"
	"zzzz/mytime"
)

type A struct {
	a [] string
}
func main() {
	mp := make(map[string]int)
	mp["hello"] = 0
	fmt.Println(len(mp))
	x := A{}
	x.a = append(x.a, "who")
	for _, v := range x.a {
		fmt.Println(v)
	}
	mytime.TestSleepRandomSecond()
}
