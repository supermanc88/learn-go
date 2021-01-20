package main

import "fmt"

// 全局作用域

var x = 100

// 100
func f1() {
	// 函数中，查找变量的顺序，
	// 1. 先在函数内部查找
	// 2. 找不到就往函数外面找，一直找到全局变量

	name := "beijing"
	// 函数内部定义的变量只能在该函数内部使用

	fmt.Println(name)
	fmt.Println(x)
}

func main() {
	f1()

	// 语句块作用域
	for i := 10; i < 18; i++ {
		fmt.Println(i)
	}

}
