package main

import (
	"fmt"
	"time"
)

func main() {
	
	now := time.Now()

	fmt.Printf("current time: %v\n", now.Unix())
	fmt.Printf("current time(nano): %v\n", now.UnixNano())

	time1 := time.Now()
	time2 := time.Now()

	if time1.Before(time2) == true {
		fmt.Println("time1 is before time2...")
	}

	// 定时器 time.Tick
	ticker := time.Tick(time.Second) // 每1秒钟执行一次time.Second
	for i := range(ticker) {
		fmt.Printf("tick.... %v\n", i)
	}
}
