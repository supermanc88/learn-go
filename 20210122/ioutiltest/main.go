package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	ret, err := ioutil.ReadFile("E:/Go/src/github.com/supermanc88/learn-go/20210122/filetest/main.go")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println(string(ret))
}
