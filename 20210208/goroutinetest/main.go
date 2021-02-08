package main

import (
	"fmt"
	"sync"
	"time"
)

// 会出现以下情况：
// 94
// 91
// 95
// 100
// 100
// 100
// 100
// 100
//
// Process finished with exit code 0

// 原因：是因为匿名函数中的i是从外面取值，当打印函数执行的时候，i可能不是执行 go func时候的值
// 如，在第55次调用go func的时候，i为55，当 匿名函数中的打印执行的时候，i可能变成58了
// 所以就打印58
func goroutineTest() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	// 使用休眠函数等待所有的goroutine执行完成
	time.Sleep(time.Second)
}

func waitByChannel() {
	var ch = make(chan int)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
			ch <- i
		}()
	}

	// 使用管道等待所有的goroutine执行完成
	for i := 0; i < 100; i++ {
		<-ch
	}
}

// 多次执行下面的代码，会发现每次打印的数字的顺序都不一致。
// 这是因为所有的goroutine是并发执行的，而goroutine的调试是随机的
func waitGroupTest() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1) // 每启动一个goroutine就登记+1
		go func() {
			defer wg.Done() // goroutine结束就登记-1
			fmt.Println(i)
		}()
	}
	wg.Wait() // 等待wg的计数器减为0
}

func main() {
	waitGroupTest()
}
