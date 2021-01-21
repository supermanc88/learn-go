package main

import "fmt"

// 嵌套结构体

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	addr address
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "chm",
		age:  111,
		addr: address{
			province: "hebei",
			city:     "baoding",
		},
	}

	fmt.Println(p1)

	fmt.Println(p1.age)
	fmt.Println(p1.addr.city)

	//fmt.Println(p1.city) // 这样访问需要使用匿名嵌套结构体 先在自己结构体中查找，找不到就去匿名嵌套的结构体中查找
}
