package main

import "fmt"

// 单向通道

// 只能发送
func f1(ch1 chan<- int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	//<-ch1 此时这个就报错了，通道不能发送数据
}

// 只能接收
func f2(ch2 <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch2)
	}
	//ch2 <- 1 此时这个就报错了，1不能发送到通道
}

func main() {

}
