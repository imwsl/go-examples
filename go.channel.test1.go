package main

import (
	"fmt"
)

// go 语言用的是CSP communicating sequential processes...
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启一个goroutine, 存入数据
	go func() {
		for i := 0; i < 100; i ++ {
			ch1 <- i
		}
		close(ch1) // 存完就关闭ch1
	}()

	go func() {
		// 无限循环取
		for {
			i, ok := <-ch1
			if !ok {
				break
			} else {
				ch2 <- i // 存ch2中
			}
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}
