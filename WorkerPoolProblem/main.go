package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int, 10)

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wg.Done()
			for job := range ch {
				fmt.Printf("Worker %d processed Job: %d: result: %d\n", i, job, job*job)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	fmt.Println("All jobs are finished")
}
