package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	var counter int

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
