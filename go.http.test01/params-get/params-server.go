package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"))
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read post err ", err)
		return
	}	
	fmt.Println("handle post ", string(b))
	answer := `{"status": "post ok"}`
	w.Write([]byte(answer))
}

func main() {

	http.HandleFunc("/get", handle)
	http.HandleFunc("/post", handlePost)
	
	fmt.Println("start :9090 ...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("start serve err : ", err)
	}
}
