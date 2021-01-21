package main

import "fmt"

// 使用值接收者和使用指针接收者的区别
// 使用值接收者实现的接口，结构体类型或其指针类型都能存
// 使用指针接收者实现的接口，只能使用结构体的指针变量类型

// 接口和类型的关系
// 多个类型可以实现同一个接口
// 一个类型也可实现多个接口
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet uint8
}

// 使用的值接收者
//func (c cat) move() {
//	fmt.Println("cat move")
//}
//
//func (c cat) eat(s string) {
//	fmt.Printf("cat eat %s\n", s)
//}

func (c *cat) move() {
	fmt.Println("cat move")
}

func (c *cat) eat(s string) {
	fmt.Printf("cat eat %s\n", s)
}

func main() {

	var a1 animal

	c1 := cat{
		name: "chm",
		feet: 1,
	}

	c2 := &cat{
		name: "xiaoxiao",
		feet: 1,
	}

	a1 = &c1
	fmt.Println(a1) // {chm 1}
	a1 = c2
	fmt.Println(a1) // &{xiaoxiao 1}

	fmt.Println(c1, c2)

}
