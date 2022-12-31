package main

import (
	"zzzz/mybasic"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	mybasic.TestStrconv()
}
