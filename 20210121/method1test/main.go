package main

// 给自己定义的类型添加方法
// 只能给自己包的类型添加方法，不能给别的包中的类型添加方法

import "fmt"

type myInt int // 自己定义一个类型

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100)
	m.hello()
}
