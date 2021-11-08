package main

import (
	"fmt"
	"time"
)

// 测试延时功能
func main() {
	
	//(1)
	time.Sleep(time.Second)

	//(2)
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("2秒到")

	//(3)
	<-time.After(2 * time.Second)
	fmt.Println("After 两秒到")

	for {}
}
