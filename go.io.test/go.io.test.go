// io
// os.Stdin 标准输入	
// os.Stdout 标准输出
// os.Stderr  标准错误
		

package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
		
	file, err := os.Open("./go.io.test.go")
	if err != nil {
		fmt.Println("Open File error : %v\n", err)
		return
	}
	defer file.Close()

	var buf [128]byte
	var content []byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		
		if err != nil {
			fmt.Println("read file err: %v\n", err)
			break
		}

		content = append(content, buf[:n]...)
	}

	fmt.Println(string(content))
}
