package myconcurrency

import (
	"sync"
	"time"
)

func WaitGroupTest() {
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(wg *sync.WaitGroup) {
			time.Sleep(3 * time.Second)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
