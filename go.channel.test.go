package main

import (
	"fmt"
)
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功： ", ret)
}

func recv1(c chan int) {
	for {
		if data, ok := <-c; ok {
			fmt.Println("接收数据2: ", data)
		} else {
			break
		}
	}
}

// go 语言用的是CSP communicating sequential processes...
func main() {
	ch := make(chan int) // 这一种属于无缓冲通道，是阻塞的
	go recv(ch) //  添加一个接收者
	ch <- 10 // 向Channel里面传递一个数据
	fmt.Println("发送成功!!")

	ch1 := make(chan int, 2)
	go recv1(ch1)
	for i := 0; i < 10; i ++ {
		ch1 <- i
	}
}
