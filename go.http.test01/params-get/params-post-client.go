package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func main() {
	
	url := "http://127.0.0.1:9090/post"
	contentType := "application/json"
	data := `{"name": "posting"}`
	
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("http.Post err ", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error!")
		return
	}
	fmt.Println(string(b))
}
