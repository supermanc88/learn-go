package main

import "fmt"

// 自定义类型和类型别名

// 自定义类型
// 在编译之后，还是我这个类型
// type 后面跟的是类型
type myInt int

// 类型别名
// 在编译之后，就变成原来的类型
// byte 和 rune也是类型别名
type yourInt = int

func main() {
	var n myInt
	n = 100
	fmt.Printf("%T\n", n) // main.myInt main包的myInt类型

	var m yourInt
	m = 100
	fmt.Printf("%T\n", m) // int

	var c rune // type rune = int32	rune是一个字符类型
	c = '中'
	fmt.Printf("%T\n", c) // int32
}
