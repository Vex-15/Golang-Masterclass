package main

import (
	"context"
	"fmt"
	"time"
)

func ping(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return

		case ch <- fmt.Sprintf("ping %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}
func pong(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return

		case ch <- fmt.Sprintf("pong %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pingCh := make(chan string)
	done := make(chan struct{})

	go ping(ctx, pingCh)
	go pong(ctx, pingCh)

	go func() {
		timeout := time.After(5 * time.Second)
		for {
			select {
			case msg := <-pingCh:
				fmt.Println(msg)
			case <-timeout:
				fmt.Println("Timeout reached")
				close(pingCh)
				done <- struct{}{}
				return
			}
		}
	}()
	done <- struct{}{}
	fmt.Println("Main function completed")

}
