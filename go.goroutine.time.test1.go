package main

import (
	"fmt"
	"time"
)

func main() {
	timer2 := time.NewTimer(time.Second)
	for {
		<-timer2.C
		fmt.Println("时间到...")
	}

	for {}
}
