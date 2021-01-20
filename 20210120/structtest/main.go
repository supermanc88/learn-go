package main

import (
	"fmt"
)

// struct
// 结构体
// 结构体是值类型的，不是引用类型的
// 类型名在同一个包内不能重复
// type 类型名 struct {
// 		字段名	类型
// 		...
// }

type persion struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 结构体实例化
	var p persion

	// 通过字段赋值
	p.name = "chm"
	p.age = 100
	p.gender = "man"
	p.hobby = []string{"soccor", "basketball"}
	fmt.Printf("%T\n", p) // main.persion
	fmt.Println(p)

	// 通过字段访问内容
	fmt.Println(p.name)
	fmt.Println(p.age)

	var p2 persion
	p2.name = "xiaoxiao"
	p2.age = 100
	fmt.Println(p2)

	// 匿名结构体，适用于临时使用
	var s struct {
		name string
		age  int
	}
	s.name = "ch"
	s.age = 100
	fmt.Println(s)
	fmt.Printf("%T %v\n", s, s) // struct { name string; age int } {ch 100}
}
