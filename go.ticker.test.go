package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0

	// 创建一个协程
	go func() {
		for {	// 协程中接收ticker数据
			i ++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
				fmt.Println("停止ticker...")
			}
		}
	}()
	for {}
}
