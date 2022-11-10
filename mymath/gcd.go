package mymath

func Mygcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Mygcd(b, a%b)
}
