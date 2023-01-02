package mytime

import (
	"fmt"
	"time"
)

func TestSelect() {
	ch := make(chan bool)
	go func() {
		x := 0
		for {
			stop := false
			select {
			case <-time.After(1 * time.Second):
				x++
				fmt.Printf("%vs passed\n", x)

			case <-ch:
				stop = true
			}
			if stop {
				break
			}
		}
	}()
	time.Sleep(10 * time.Second)
	ch <- true
}
