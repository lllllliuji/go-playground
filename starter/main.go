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
	fmt.Println(mymath.Mygcd(10, 5))
	array := [...]int{10, 20, 30, 40} // array
	slice := []int{10, 20, 30, 40}    // slice
	for _, v := range array {
		fmt.Println(v)
	}
	for _, v := range slice {
		fmt.Println(v)
	}
	dict := make(map[string]int)
	dict["hello"] = 2
	TestType()
	TestChan()
	TestBuffedChan()
}
