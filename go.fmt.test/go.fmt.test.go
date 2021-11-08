// 学习fmt函数
//
//

/// func Print(a ...interface{}) (n int, err error)
/// 直接输出

/// func Printf(format string, a ...interface{})(n int, err error)
/// 输出带格式的字符串

/// func Println(a ...interface{})(n int, err error)
/// 输出的内容后面带一个换行符

package main 

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hello 世界!")
	name := "高铁"
	fmt.Printf("%s\n", name)
	fmt.Println("Hello World!!")

/// Fprint 是将内容输入到io.Writer接口
// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
// 直接输入内容到io.Writer接口
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// 带格式的输出
// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
// 带换行符的输出

	fmt.Fprintln(os.Stdout, "向标准输出写入内容!!")

	// 打开一个文件
	file, err := os.OpenFile("xx.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错err: ", err)
		return 
	}

	fmt.Fprint(file, "直接写入换行")
	name1 := "高铁南"
	fmt.Fprintf(file, "%s\n", name1)
	fmt.Fprintln(file, "最后一行...")

/// func Sprint(a ...interface{}) string
/// 返回一个字符串
/// func Sprintf(format string, a ...interface{}) string
/// 返回一个带格式的字符串
/// func Sprintln(a ...interface{}) string
/// 返回一个换行的字符串

	a := fmt.Sprint(22)
	b := fmt.Sprintf("name: %s  age: %d\n", "Jack", 29)
	c := fmt.Sprintln("Hello World!!")
	fmt.Println(a, b, c)

/// func Errorf(fotmat string, a ...interface{}) error
/// 生成一个格式化的error
/// 通常用来定义一个错误类型
	err = fmt.Errorf("error code : %d", 33)
	fmt.Println(err)
}