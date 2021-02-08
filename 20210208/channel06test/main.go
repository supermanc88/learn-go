package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
// 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
// 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 主goroutine从resultChan取出结果并打印到终端输出

func genRand(ch chan<- int) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Int()
	fmt.Println(x)
	ch <- x
	time.Sleep(time.Millisecond * 500)
}

func addSum(ch <-chan int, result chan<- int) {
	for x := range ch {
		fmt.Println(x)
		var sum int
		for x != 0 {
			sum += x % 10
			x /= 10
		}
		fmt.Println(sum)
		result <- sum
	}

}

func main() {
	jobChan := make(chan int, 100)
	resultChan := make(chan int, 100)

	for w := 0; w < 24; w++ {
		go addSum(jobChan, resultChan)
	}

	for i := 0; i < 100; i++ {
		genRand(jobChan)
	}
	close(jobChan)

	for i := 0; i < 100; i++ {
		<-resultChan
	}

}
