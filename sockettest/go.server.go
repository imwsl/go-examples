package main

import (
	"fmt"
	"net"
	"bufio"
	"io"

	"goexamples/sockettest/proto"
)

// 处理socket链接
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("读取数据错误: ", err)
			return 
		}
		fmt.Println("获得服务端的数据: ", msg, "\n")
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("监听错误!, err: ", err)
		return 
	}
	fmt.Println("开启TCP:127.0.0.1:20000...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("建立连接错误!, err: ", err)
			continue
		}

		// 开启一个goroutine处理链接
		go process(conn)
	}
}