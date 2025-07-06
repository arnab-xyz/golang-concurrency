package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(wg *sync.WaitGroup, ch1 chan<- string, ctx context.Context) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("Shutting down worker...")
	case <-time.After(time.Second * 2):
		ch1 <- "Hello..."
	}
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go Worker(&wg, ch1, ctx)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout occurred...")
		cancel()
	}
	wg.Wait()
	fmt.Println("Finishing...")
}
