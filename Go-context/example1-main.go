package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Create a context with a timeout of 3 seconds
	ctxWithTimeout, cancelTimeout := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelTimeout()

	// Create a context with a deadline 5 seconds from now
	deadline := time.Now().Add(5 * time.Second)
	ctxWithDeadline, cancelDeadline := context.WithDeadline(context.Background(), deadline)
	defer cancelDeadline()

	wg.Add(1)
	go func() {
		defer wg.Done()
		childProcess(ctxWithTimeout, "プロセス1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		childProcess(ctxWithDeadline, "プロセス2")
	}()

	wg.Wait()
}

func childProcess(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			// Handle cancellation or timeout
			fmt.Printf("%s canceled or timed out. error: %s\n", name, ctx.Err())
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("%s processing...\n", name)
		}
	}
}
