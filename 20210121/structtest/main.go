package main

import "fmt"

// struct
// 构造函数

type person struct {
	name string
	age  int
}

// 构造函数	约定俗成 用new开头
// 返回的是结构体还是结构体指针
// 当结构体比较大的时，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("chm", 100)
	p2 := newPerson("xiaoxiao", 100)

	fmt.Println(p1, p2)
}
