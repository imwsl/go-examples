//cpy_it.go
//

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	
	if len(os.Args) != 3 {
		fmt.Println("command lines error!\n")
		return 
	}

	fromFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("open file error: %v\n", err)
		return
	}

	destFile, err1 := os.Create(os.Args[2])
	if err1 != nil {
		fmt.Println("create file error: %v\n", err1)
		return
	}

	defer fromFile.Close()
	defer destFile.Close()

	var buf [1024]byte
	for {
		n, err := fromFile.Read(buf[:])
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Read from file error: %v\n", err)
			break
		}

		destFile.Write(buf[:n])
	}

	fmt.Println("-----------------THE END......................")
}
