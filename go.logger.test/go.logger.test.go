//
// log: golang provides this library to output the logs more elegant!!
//

package main

import (
	"log"
	"os"
	"fmt"
)

func init() {
	logFile, err := os.OpenFile("./xxx.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Open log file error: %v\n", err)
		return
	}

	log.SetOutput(logFile)
}

func main() {

	// setFlags
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	
	log.Println("This is a normal log....")

	v := "Error Found!!"

	log.Printf("This is a %s\n", v)

	log.SetPrefix("[fatal error]")
	log.Fatalln("This is a fatal log")
	log.Panicln("This is a panic log")

	// log.New
	logsFile, err := os.OpenFile("./log.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("new os.OpenFile err %v\n", err)
		return
	}

	logger := log.New(logsFile, "<new>", log.Lshortfile | log.Ldate | log.Ltime)
	logger.Println("Simple Logger....")
	fmt.Println("end...")
}
