lspackage main

import (
	"fmt"
	"time"
)

// 重置定时器
func main() {
	timer := time.NewTimer(3 * time.Second)
	//timer.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer.C)

	for {}
}
