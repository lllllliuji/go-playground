package main

import (
	"fmt"
	"zzzz/mymath"
)

func init() {
	fmt.Println("init")
}

func main() {
	str := "world"
	fmt.Println("hello", str)
	fmt.Println(mymath.Sqrt(100.0))
}
