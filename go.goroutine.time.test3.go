package main

import (
	"fmt"
	"time"
)

// 测试延时功能
func main() {
	timer := time.NewTimer(2 * time.Second)
	go func(){
		<-timer.C
		fmt.Println("定时器执行了...")
	}()
	b := timer.Stop()
	if b { // 定时器关闭，上面的go func不会执行
		fmt.Println("定时器已关闭!!")
	}

	for {}
}
