package main

import "fmt"

// go中推荐使用驼峰命名
//var name string
//var age int
//var isOK bool

// 全局变量不使用是可以编译过去的
// 局部变量是不可以的
var (
	name string
	age int
	isOK bool
)

func main() {
	name = "chm"
	age = 28
	isOK = true

	// 变量声明了必须使用，不使用编译不过去
	fmt.Printf("name : %s", name) // 字符串格式化
	fmt.Println()
	fmt.Println(age) // 加换行
	fmt.Print(isOK) // 输出 没换行
	fmt.Println()

	// 声明变量同时赋值 同一个作用域不能重复声明一个变量
	var s1 string  = "aaa"
	fmt.Println(s1)

	// 类型推导 根据值推断类型
	var s2 = 20
	fmt.Println(s2)

	// 短变量声明 只能在函数中使用 相当于上面的简写
	s3 := "hahaha"
	fmt.Println(s3)

	// 匿名变量
	// 有的时候不使用这个变量，就使用 _ 接收 多用于占位，表示忽略，一般用于忽略返回值
}
