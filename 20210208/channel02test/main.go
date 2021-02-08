package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel练习
// 1. 启动一个goroutine，生成100个数发送到ch1
// 2. 启动另一个goroutine，从ch1中取值，计算其平方并放到ch2
// 3. 在main中 从 ch2 取值并打印

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {

		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 10; i++ {
			ch1 <- rand.Intn(10)
		}
	}()

	go func() {

		for i := 0; i < 10; i++ {
			x := <-ch1
			ch2 <- x * x
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch2)
	}
}
