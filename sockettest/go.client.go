package main

import (
	"fmt"
	"net"

	"goexamples/sockettest/proto"
)



func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接tcp错误: ", err)
		return 
	}

	defer conn.Close()  // 最后关闭连接
	for i := 0; i < 20; i ++ {
		msg := "Hello, Hello. How are you?"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed: err: ", err)
			return
		}
		conn.Write(data)
	}
}