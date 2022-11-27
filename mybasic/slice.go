package mybasic

import "fmt"

func TestSlice() {
	primes := []int{2, 3, 5, 7, 11, 13}
	p := primes[1:]
	fmt.Println(p)
	var sli [] int
	sli = append(sli, primes[4:]...)
	fmt.Println(sli)
	primes = primes[0:4]
	fmt.Println(primes)
}
