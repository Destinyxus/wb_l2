package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		var wg sync.WaitGroup
		wg.Add(len(channels))

		// For each input channel, start a goroutine to wait for its closure
		for _, c := range channels {
			go func(c <-chan interface{}) {
				defer wg.Done()
				select {
				case <-c:
					select {
					case <-orDone:
					default:
						close(orDone)
					}
				case <-orDone:
				}
			}(c)
		}

		// Wait for all goroutines to finish
		wg.Wait()
	}()

	return orDone
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
