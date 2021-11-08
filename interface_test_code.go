package main

import (
	"fmt"
)

func main() {
	
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "Alice"
	studentInfo["age"] = 18
	fmt.Println(studentInfo)

	var x interface{}
	x = "Hello World"	
	v, ok := x.(string) 

	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("x 不是 string...")
	}
}
