package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int = 0
	var mutex sync.Mutex
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
