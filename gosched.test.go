package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 开一个协程
	go func(s string) {
		for i := 0; i < 2; i ++ {
			fmt.Println(s)
		}
	}("协程")

	// 主协程，main开始会开一个主协程
	for i := 0; i < 2; i ++ {
		runtime.Gosched()
		fmt.Println("主协程")
	}
}
