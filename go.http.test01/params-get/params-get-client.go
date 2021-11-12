package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
)

func main() {
	
	apiUrl := "http://127.0.0.1:9090/get"

	data := url.Values{}
	data.Set("name", "JSON")
	
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("ParseRequestURI err ", err)
		return
	}

	u.RawQuery = data.Encode()
	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("http.Get err ", err)
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
