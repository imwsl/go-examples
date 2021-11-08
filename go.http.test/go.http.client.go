// http客户端
// go.http.client.go
//

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, _ := http.Get("http://127.0.0.1:9000/go")
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("获取错误: ", err)
			return 
		} else {
			fmt.Println("读取完毕...")
			fmt.Println(string(buf[:n]))
			break
		}
	}
}