package main

import "fmt"

// 函数类型
// 可以作为参数，也可以作为返回值

// 匿名函数就是没有名字的函数

func f1() {
	fmt.Println("hello world")
}

func f2() int {
	return 10
}

// 就相当于函数指针
func f3(x func()) {
	x()
}

func f4(x, y int) int {
	return x + y
}

// 就是套娃
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	a := f1
	fmt.Printf("%T\n", a) // func()
	b := f2
	fmt.Printf("%T\n", b)  // func() int
	fmt.Printf("%T\n", f4) // func(int, int) int

	f3(f1)

	f7 := f5(f2)
	fmt.Println(f7)
}
