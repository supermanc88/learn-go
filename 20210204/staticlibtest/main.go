package main

import (
	"fmt"
	"github.com/supermanc88/btcnet"
)

// 使用静态库
// "github.com/supermanc88/" 这个路径是pkg的相对路径 btcnet 是库的名字
// 使用方式：
// 编译：
// go tool compile -I E:\Go\pkg\windows_amd64 .\main.go
// 链接：
// go tool link -o main.exe -L E:\Go\pkg\windows_amd64\ main.o

func main() {

	data := btcnet.Sm3Sum256([]byte("teststring"))

	fmt.Println(data)

}
