package main

import "fmt"

// panic recover

func f1() {
	fmt.Println("A")
}

func f2() {
	// 刚刚打开了数据库连接
	defer func() {
		err := recover() // 尽量少用 必须搭配defer使用，defer一定要在可能panic的语句之前使用
		fmt.Println(err)
		fmt.Println("释放数据库数据")
	}()
	panic("出现错误了") // 造成程序崩溃退出 后面就不执行了
	fmt.Println("B")
}
func f3() {
	fmt.Println("C")
}

func main() {

	// A
	// 出现错误了
	// 释放数据库数据
	// C
	f1()
	f2()
	f3()
}
