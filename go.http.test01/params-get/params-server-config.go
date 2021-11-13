package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"log"
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

type MyHandler struct {}
func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello ServeHTTP"))
}

func main() {

	http.HandleFunc("/get", handle)
	http.HandleFunc("/post", handlePost)
	
	handler := MyHandler{}
	
	s := &http.Server{
		Addr:			":9090",
		Handler: 	handler,
		ReadTimeout:	10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}	
	fmt.Println("start :9090 ...")
	log.Fatal(s.ListenAndServe())
}
