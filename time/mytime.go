package mytime

import (
	"fmt"
	"time"
)

func TestTime() {
	t := time.Now()
	u := t.Add(-1 * time.Hour)
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", u)
}