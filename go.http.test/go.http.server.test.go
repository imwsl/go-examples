// 出差梧州，还是喜欢写几行代码，学习net/http
// go.http.server.test.go
//

package main

import (
	"fmt"
	"net/http"
)

func main() {

	// 开启http://127.0.0.1:9000/go 服务
	http.HandleFunc("/go", myHandler)

	// 设置监听
	fmt.Println("开启127.0.0.1:9000...")
	http.ListenAndServe("127.0.0.1:9000", nil)
}

// http相应函数
func myHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.RemoteAddr, "连接成功!")
	fmt.Println("method: ", r.Method)
	fmt.Println("url: ", r.URL.Path)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)

	w.Write([]byte("Hello World!!!"))
}