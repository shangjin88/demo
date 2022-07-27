package main

import (
	"errors"
	"fmt"
	"github.com/avast/retry-go"
	"time"
)

func main() {

	err := retry.Do(
		func() error {
			return errors.New("test")
		},
		retry.Attempts(3),
		retry.Delay(100*time.Millisecond),
		// 只拿最后一个错误
		retry.LastErrorOnly(true),
		// 重试失败的回调函数
		retry.OnRetry(func(n uint, err error) {
			fmt.Println("OnRetry...")
		}),
		// 满足条件才允许重试
		retry.RetryIf(func(error) bool {
			return false
		}),
	)

	if err != nil {
		fmt.Println(err)
	}

}
