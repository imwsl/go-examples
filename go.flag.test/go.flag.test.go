// flag: 主要功能为定义命令行参数
//

package main

import (
	"fmt"
	"flag"
	"time"
)

func main() {
	var name string
	var age int
	var married bool
	var delay time.Duration

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "duration", 0, "延迟的时间间隔")

	// parse command line
	flag.Parse()
	fmt.Println(name, age, married, delay)
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}


