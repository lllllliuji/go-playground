package mymath

import "fmt"

func Sqrt(x float64) float64 {
	result := x
	for i := 0; i < 10; i++ {
		if result <= 0 {
			result = 0.1
		}
		delta := x - (result * result)
		result = result + 0.5*delta/result
		fmt.Println("comuputing sqrt of", x, "to be", result)
	}
	return result
}
