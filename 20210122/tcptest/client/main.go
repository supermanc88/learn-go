package main

import (
	"fmt"
	"net"
)

// client

func main() {
	// 1. 与server建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial failed, err : %v\n", err)
		return
	}
	// 2. 发送数据
	conn.Write([]byte("test test tses"))

	conn.Close()
}
