package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {

	for {

		once.Do(func(){
			fmt.Println("For循环中只执行一次...")
		})

		fmt.Println("For循环...")
		time.Sleep(1 * time.Second)
	}
}
