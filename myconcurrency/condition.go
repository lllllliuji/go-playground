package myconcurrency

import (
	"fmt"
	"sync"
)

func Condition_test() {
	var mut sync.Mutex
	cond := sync.NewCond(&mut)
	count := 0
	for i := 0; i < 10; i++ {
		go func(x int) {
			mut.Lock()
			defer mut.Unlock()
			fmt.Println(x)
			count++
			cond.Broadcast()
		}(i)
	}

	mut.Lock()
	for count < 10 {
		cond.Wait()
	}
	mut.Unlock()
}
