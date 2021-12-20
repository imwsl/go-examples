//
// go.format.test.go
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
	Name string
	Hobby string
}

func main() {
	p := Person{"Wesley", "Finishing"}

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
}

