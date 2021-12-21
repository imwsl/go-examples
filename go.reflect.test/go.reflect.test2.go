// 
//
package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Id int `json: "id" db:"index_id"`
	Name string
	Age int
}

type Dog struct {
	Animal
	Color string
}

func (a Animal) Hello() {
	fmt.Println("say Hello")
}

func reflect_output(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t)
	fmt.Println("Name: ", t.Name())	
	
	ft := t.Field(0).Tag
	fmt.Println("tag:\n")
	fmt.Println(ft.Get("json"))


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

	ttt := reflect.TypeOf(a)
	ftt := ttt.Field(0).Tag
	fmt.Printf("ft.tag %s\n", ftt.Get("db"))

	reflect_output(a)

	m := Dog{Animal{2, "John", 22}, "red"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	fmt.Printf("%#v  %#v\n", t.Field(0))
	fmt.Printf("%#v \n", reflect.ValueOf(m).Field(0))
	
	for i := 0; i < t.NumField(); i ++ {
		f := t.Field(i)
		fmt.Printf("\nt: %#v %#v\n", f.Name, f.Type);
	
		if f.Name == "Animal" {
			val, _ := reflect.ValueOf(m).Field(i).Interface().(Animal)
			fmt.Printf("\n\n id = %v name = %v Age = %v \n\n", val.Id, val.Name, val.Age)
		}
	}
	
	tt := reflect.ValueOf(m)
	for i := 0; i < tt.NumField(); i ++ {
		ff := tt.Field(i)
		fmt.Printf("\ntt: %#v\n", ff)
	}
}

