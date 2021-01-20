package main

import (
	"fmt"
)

// 匿名函数

func main() {

	// 在函数内部没有办法声明带名字的函数，只能定义匿名函数
	var f1 = func(x, y int) {
		fmt.Println(x + y)
	}

	f1(10, 20)

	// 如果只调用一次的话，还可以简写成立即执行
	func() {
		fmt.Println("hello world")
	}()

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	f2 := func(x, y int) int {
		return x + y
	}(10, 20)
	fmt.Println(f2)
}
