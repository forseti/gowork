package helper

import (
	"context"
	"time"
)

func CreateCounter(ctx context.Context, lag int64) chan int {
	dest := make(chan int)

	go func() {
		defer close(dest)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dest <- counter
				counter++
				time.Sleep(time.Duration(lag) * time.Second)
			}
		}
	}()

	return dest
}
