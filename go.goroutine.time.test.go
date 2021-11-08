package main

import (
	"fmt"
	"time"
)

func main() {
	
	timer1 := time.NewTimer(2 * time.Second) // 2秒
	t1 := time.Now()
	fmt.Printf("t1: %v\n", t1)
	t2 := <-timer1.C  // 阻塞等待定时器
	fmt.Printf("t2:%v\n", t2)

	for {}
}
