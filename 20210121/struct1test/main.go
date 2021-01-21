package main

import "fmt"

// 结构体中会出现的一些问题

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 结构体的匿名字段
// 字段比较少且简单的
type dog struct {
	string
	int
}

func main() {
	// 方法1
	var p person
	p.name = "chm"
	p.age = 1

	// 方法2
	p1 := person{
		name: "xiao",
		age:  2,
	}
	fmt.Println(p1)

	// 方法3
	p2 := person{
		"xiao",
		3,
	}
	fmt.Println(p2)

	d := dog{
		string: "aaa",
		int:    0,
	}
	// 把类型当成名字使用了
	fmt.Println(d.string, d.int)
}
