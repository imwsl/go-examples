package main 

import (
	"fmt"
)

func main() {
	
	// %v 默认格式
	// %+%v 但部分情况与%v相同，输出结构体的时候会添加字段名
	// %#v 值的go语法表示
	// %T 打印值得类型
	// %% 百分号
	
	o := struct{name string} {"Hello World"}
	
	fmt.Printf("直接输出: %v\n", "Hello World")
	fmt.Printf("输出结构体显示字段: %+v\n", o)
	fmt.Printf("go语法表示: %#v\n", o)
	fmt.Printf("打印值类型: %T\n", o)
	fmt.Printf("100%%\n") 

	// %t 输出true or false
	// %b 输出二进制
	// %c 输出该值得uncode码值
	// %d 十进制
	// %O 八进制
	// %x 十六进制 a-f
	// %X 十六进制 A-F
	// %U 输出Unicode格式
	// %q 单引号括起来的go语法字面值

	n := 65
	fmt.Printf("%t\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%O\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X", n)
}