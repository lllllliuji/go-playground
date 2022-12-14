package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	s := []int{}
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	s1 := s[10: 20]
	s1 = append(s1, 100)
	s[10] = 99
	s[11] = 98
	s[12] = 97
	fmt.Println(s1)
	index := 0
	for index < len(s) {
		s2 := s[index:min(index+10, len(s))]
		fmt.Println(s2)
		index = min(index+10, len(s))
	}
}
