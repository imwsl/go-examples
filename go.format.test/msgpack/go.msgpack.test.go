//msgpack
//

package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"io/ioutil"
	"math/rand"
)

type Person struct {
	Name string
	Age int
	Sex string
}

func writeJson(filename string) (err error) {
	var persons []*Person

	for i := 0; i < 10; i ++ {
		p := &Person{
			Name: fmt.Sprintf("name:%d", i),
			Age: rand.Intn(100),
			Sex: "male",
		}

		persons = append(persons, p)
	}

	// binary json
	data, err := msgpack.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	return
}

func readJson(filename string) (err error) {
	var persons []*Person

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return 
	}

	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range persons {
		fmt.Println("%#v\n", v)
	}

	return
}

func main() {
	
	writeJson("./my.dat")

	readJson("./my.dat")
}


