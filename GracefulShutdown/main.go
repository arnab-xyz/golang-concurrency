package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Shutting down gracefully...")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go Worker(&wg, ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("Finishing...")
}
