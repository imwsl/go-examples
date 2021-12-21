// 
//
package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Id int
	Name string
	Age int
}

func (a Animal) Hello() {
	fmt.Println("say Hello")
}

func reflect_output(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t)
	fmt.Println("Name: ", t.Name())	
	

	// field: name & type
	//
	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i ++ {
		f := t.Field(i)
		fmt.Printf("field: %s - %v\n", f.Name, f.Type)
		
		val := v.Field(i).Interface()
		fmt.Println("val: ", val)
	}

	for i := 0; i < t.NumMethod(); i ++ {
		m := t.Method(i)
		fmt.Printf("method: %s - %v\n", m.Name, m.Type)
	}
}

func main() {
	a := Animal{1, "Joe", 19}
	reflect_output(a)
}

