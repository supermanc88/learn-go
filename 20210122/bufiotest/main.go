package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fileObj, err := os.Open("E:/Go/src/github.com/supermanc88/learn-go/20210122/filetest/main.go")
	if err != nil {
		fmt.Printf("open file failed, err :%v\n", err)
		return
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("eof")
			return
		}
		if err != nil {
			fmt.Printf("read line failed error: %v", err)
			return
		}

		fmt.Println(str)
	}
}
