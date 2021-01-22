package main

import (
	"fmt"
	"net"
)

// tcp server

func main() {
	// 1. 本地端口启动
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("start server failed, err = %v\n", err)
		return
	}
	// 2. 等待别人跟我建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("accept failed, err : %v\n", err)
		return
	}
	// 3. 与客户端通信
	var temp [128]byte
	n, err := conn.Read(temp[:])
	if err != nil {
		fmt.Printf("read from conn failed, err : %v\n", err)
		return
	}
	fmt.Println(string(temp[:n]))

}
