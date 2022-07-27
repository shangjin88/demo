package main

import (
	"context"
	"fmt"
	"time"
)

func task() error {
	return fmt.Errorf("task error")
}

func do() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second)

	defer cancel()

	for {
		if err := task(); err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("context timeout")
		case <-time.After(5 * time.Second):
			return fmt.Errorf("task timeout")
		}
	}
}

func main() {
	err := do()
	fmt.Println(err)
}
