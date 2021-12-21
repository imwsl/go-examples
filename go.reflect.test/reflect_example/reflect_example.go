package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Server struct {
	IP string
	Port string
}

type MySQL struct {
	UserName string
	Password string
	Database string
	Host string
	Port string
	Timeout string
}

func is_file_existed(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func main() {
	cfg, err := ini.Load("./my.ini")
	if err != nil {
		fmt.Println("load ini err: ", err)
		return
	}

	ip := cfg.Section("server").Key("ip").String()
	port := cfg.Section("server").Key("port").String()

	fmt.Println("ip => ", ip, " port => ", port)

	s := Server{ip, port}

	username := cfg.Section("mysql").Key("username").String()
	passwd := cfg.Section("mysql").Key("passwd").String()
	database := cfg.Section("mysql").Key("database").String()
	host := cfg.Section("mysql").Key("host").String()
	port = cfg.Section("mysql").Key("port").String()
	timeout := cfg.Section("mysql").Key("timeout").String()

	m := MySQL{username, passwd, database, host, port, timeout}

	// marshal
	sbytes, _ := json.Marshal(s)
	mbytes, _ := json.Marshal(m)

	// save to the disk
	//
	server_file := "./server.dat"
	mysql_file := "./mysql.dat"
	if is_file_existed(server_file) == true {
		os.Remove(server_file)
	}

	fp1, _ := os.Create(server_file)
	defer fp1.Close()

	if is_file_existed(mysql_file) == true {
		os.Remove(mysql_file)
	}

	fp2, _ := os.Create(mysql_file)
	defer fp2.Close()

	ioutil.WriteFile(server_file, sbytes, 0777)
	ioutil.WriteFile(mysql_file, mbytes, 0777)

	rsbytes, _ := ioutil.ReadFile(server_file)
	rmbytes, _ := ioutil.ReadFile(mysql_file)

	var ss Server
	var mm MySQL
	json.Unmarshal(rsbytes, &ss)
	json.Unmarshal(rmbytes, &mm)
	
	fmt.Println("\nUnmarshal....\n")
	fmt.Println(ss)
	fmt.Println(mm)
}
