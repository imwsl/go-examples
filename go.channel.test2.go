package main

import (
	"fmt"
)

func send(ch chan<- int) {
	for i := 0; i < 100; i ++ {
		ch <- i
	}
	close(ch)
}

func square(out chan<- int, in <-chan int) {
	for {
		i, ok := <-in
		if !ok {
			break
		} else {
			out <- i * i
		}
	}
	close(out)
}

// go 语言用的是CSP communicating sequential processes...
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go send(ch1)
	go square(ch2, ch1)

	for i := range ch2 {
		fmt.Println("最后得到的数据：", i)
	}
}
