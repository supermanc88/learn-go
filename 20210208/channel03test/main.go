package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	close(ch1)

	<-ch1
	<-ch1

	x, ok := <-ch1
	fmt.Println(x, ok) // 0 false 从一个关闭的channel中依然可以取值，返回零值
}
