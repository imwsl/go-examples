package main

import (
	"fmt"
	//"time"
)

func test1(ch1 chan string) {
	for {
		ch1 <- "来自Channel 1..."
	}
}

func test2(ch2 chan string) {
	for {
		ch2 <- "来自Channel 2..."
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go test1(ch1)
	go test2(ch2)
	for {
		
		select {
			case s1 := <-ch1:
				fmt.Println("s1: ", s1)
			case s2 := <-ch2:
				fmt.Println("s2: ", s2)
		}
	}
}
