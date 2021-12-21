//
//

package main

import (
	"fmt"
	"reflect"
)

func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("Type is : ", t)
	
	k := t.Kind()
	fmt.Println(k)

	switch k {
		
		case reflect.Float64:
			fmt.Println("a is float64\n")
		case reflect.String:
			fmt.Println("a is string\n")
	}
}

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	
	switch k {
		case reflect.Float64:
			fmt.Println("a => ", v.Float())
		case reflect.String:
			fmt.Println("a => ", v.String())
	}
}

func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
		case reflect.Float64:
			v.SetFloat(8.9)
		case reflect.Ptr: // modify reflect value, must pass the pointer...
			v.Elem().SetFloat(9.9)
	}
}

func main() {
	var x float64 = 3.4
	reflect_type(x)
	reflect_value(x)
	reflect_set_value(&x)
	fmt.Println(x)
	fmt.Println("\n")

	var a string = "Hello World!!"
	reflect_type(a)
	reflect_value(a)
}
