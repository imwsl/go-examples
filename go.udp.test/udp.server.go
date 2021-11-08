// UDP服务端
//

package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建Listen
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("UDP创建Listen失败: ", err)
		return 
	}
	defer listen.Close() // 关闭UDP监听

	// 开启UDP接收数据
	fmt.Println("开启UDP服务端 0.0.0.0:30000...")
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("UDP接收数据失败: ", err)
			continue
		}
		fmt.Printf("data: %v addr: %v count: %v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("UDP发送数据失败: ", err)
			continue
		}
	}
}