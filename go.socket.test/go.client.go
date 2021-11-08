package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)



func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接tcp错误: ", err)
		return 
	}

	defer conn.Close()  // 最后关闭连接

	// 与键盘输入交互
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			fmt.Println("退出了...")
			return
		}

		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("写入错误: ", err)
			return 
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("读取错误: ", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}