package main

import "fmt"

// 闭包
// 闭包是一个函数，这个函数包含他外部作用域的变量，如adder 中的 匿名函数中使用了外部的x

// 闭包 = 函数  +  外部变量的引用

// 底层原理：
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序

// 比如别人的接口定义如f1
// 现在想把f2当参数传进去执行
func f1(f func()) {
	fmt.Println("f1")
	f()
}

func f2(x, y int) {
	fmt.Println("f2")
	fmt.Println(x + y)
}

// 定义一个函数对f2进行包装
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		fmt.Println("hello f3")
		f(x, y)
	}
	return tmp
}

// 例子2
func adder() func(int) int {
	x := 100
	return func(y int) int {
		x += y
		return x
	}
}

func main() {

	// 如何让f2传到f1中执行
	//f1(f2) 直接这样是不行的，因为类型不匹配

	f1(f3(f2, 1, 2))

	retFunc := adder()
	ret := retFunc(200)
	fmt.Println(ret) // 300

	ret = retFunc(300)
	fmt.Println(ret) // 600
}
