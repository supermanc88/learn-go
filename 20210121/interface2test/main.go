package main

import "fmt"

// 同一个结构体可以实现多个接口

// 接口还可以嵌套
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat()
}

type cat struct {
	name string
	feet uint8
}

func (c *cat) move() {
	fmt.Println("cat move")
}

func (c *cat) eat() {
	fmt.Println("cat eat")
}

func main() {
	var (
		m1 mover
		e1 eater
	)

	m1 = &cat{
		name: "chm",
		feet: 0,
	}

	e1 = &cat{
		name: "xiao",
		feet: 0,
	}

	fmt.Println(m1)
	fmt.Println(e1)
}
