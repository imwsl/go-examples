//
// go.format.test2.go
//
//

// json
// encoding/json
// func Marshal(v interface{})([]byte, error)

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Hobby string `json:"hobby"`
	Age int `json:"age"`
}

func main() {
	p := Person{"Wesley", "Finishing", 19}

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err: ", err)
		return	
	}

	fmt.Println(string(b))

	// format string to output
	//
	b, err = json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println("json format err: ", err)
		return
	}

	fmt.Println(string(b))

	// unmarshal
	// 
	from := []byte(`{"name": "Joe", "hobby": "finishing", "age": 22}`)
	var to Person
	err1 := json.Unmarshal(from, &to)
	if err1 != nil {
		fmt.Println("unmarshal err1 : ", err1)
		return 
	}
	fmt.Println(to)
}

