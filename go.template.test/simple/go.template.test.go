// template 
//

package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("create template failed. err => ", err)
		return
	}

	tmpl.Execute(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloFunc)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err: ", err)
		return
	}
}
