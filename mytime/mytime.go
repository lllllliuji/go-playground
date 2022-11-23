package mytime

import (
	"fmt"
	"math/rand"
	"time"
)

func TestTime() {
	t := time.Now()
	u := t.Add(-1 * time.Hour)
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", u)
}

func TestSleepRandomSecond() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(100) + 100) * time.Millisecond)
	fmt.Printf("Sleep %v millisecond\n", time.Since(start).Milliseconds())
}