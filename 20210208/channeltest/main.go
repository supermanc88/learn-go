package main

import (
	"fmt"
	"sync"
)

// channel
// 通道
// var 变量 chan 类型
// channel 是 引用类型
// 切片、map、channel是引用类型
// 必须初始化才能使用
// <- 表示数据的流向，没有向右的符号
// 发送 ch <- x
// 接收 x := <- ch1
// 丢弃 <- ch
// 关闭 close

var ch chan int

func noBufferedChannel() {
	ch = make(chan int)       // 无缓冲区的通道
	ch1 := make(chan int, 16) // 带缓冲区的通道

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("后台goroutine接收")
		fmt.Println(<-ch)
	}()

	ch <- 10

	fmt.Println(ch, ch1)

	wg.Wait()
}

func bufferedChannel() {
	ch1 := make(chan int, 1) // 带缓冲区的通道
	defer close(ch1)         // 关闭已经关闭的通道也会引起panic

	ch1 <- 10 // 这种的不阻塞

	// 当通道中的元素个数满之后，再向通道中传值时，会阻塞报错
	// fatal error: all goroutines are asleep - deadlock!
	// ch1 <- 20

	fmt.Println(ch1)

	fmt.Println(<-ch1)
}

func main() {
	bufferedChannel()
}
