// template 
//

package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct {
	Name string
	Age int
	Sex string
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("create template failed. err => ", err)
		return
	}

	user := User{
		Name: "Jason",
		Age: 19,
		Sex: "Male",	
	}

	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", helloFunc)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err: ", err)
		return
	}
}
