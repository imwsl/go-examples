package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	
}
func main() {

	http.HandleFunc("/get", handle)
	
	fmt.Println("start :9090 ...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("start serve err : ", err)
	}
}
