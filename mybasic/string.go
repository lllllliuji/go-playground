package mybasic

import (
	"fmt"
	"strconv"
)

func TestStrconv() {
	num, _ := strconv.Atoi("11")
	fmt.Println(num)

	str := strconv.Itoa(num * 100)
	fmt.Println(str)
}