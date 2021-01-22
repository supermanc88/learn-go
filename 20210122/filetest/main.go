package main

import (
	"fmt"
	"io"
	"os"
)

// 文件操作

func main() {
	fileObj, err := os.Open("E:/Go/src/github.com/supermanc88/learn-go/20210122/filetest/main.go")
	if err != nil {
		fmt.Printf("open file failed, err :%v\n", err)
		return
	}

	// 记得关闭文件
	defer fileObj.Close()

	// 读文件

	var data = make([]byte, 128)
	for {
		n, err := fileObj.Read(data)
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read file failed, err : %v\n", err)
			return
		}
		fmt.Printf("read %d bytes\n", n)
		fmt.Println(string(data[:n]))
		if n < 128 {
			return
		}
	}
}
