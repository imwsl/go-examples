package main

import (
	"fmt"
	"net"
	"bufio"
)

// 处理socket链接
func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("读取客户端数据错误! err: ", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到客户端发来得数据: ", recvStr)
		recvStr = "recv: " + recvStr
		conn.Write([]byte(recvStr))
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