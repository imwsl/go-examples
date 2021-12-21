package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	ID int `json:"ID" db:"INDEX_ID"`
	Name string `json:"name"`
	Age int
}

func main() {
	p := Person{19, "Joe", 22}

	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i ++ {
		f := t.Field(i).Tag
		fmt.Println("\nTag:")
		fmt.Println(f.Get("json"))
	}
}
