package main

import (
	"fmt"
	"time"
)

// goroutine
// 并发：同一时间段内执行多个任务
// 并行：同一时刻执行多个任务
// go中使用goroutine实现
// goroutine类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个goroutine并发工作。
// goroutine是由go语言的运行时runtime调度完成，而线程是由操作系统调度完成
// go语言提供channel在多个goroutine间进行通信
// goroutine和channel是go语言秉承的CSP(communicating sequential process)并发模式的重要实现基础

// go中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine
// 一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数

func hello(i int) {
	fmt.Printf("hello %d\n", i)
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 10; i++ {
		go hello(i) // 开启一个单独的goroutine去执行这个函数

		go func(i int) { // 匿名函数
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	// main函数结束之后，由main函数启动的goroutine也都结束了
}
