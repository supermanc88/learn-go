package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet uint8
}

type chicken struct {
	feet uint8
}

func (c chicken) move() {
	fmt.Println("鸡动")
}

func (c chicken) eat(s string) {
	fmt.Println("鸡吃饮料")
}

func (c cat) move() {
	fmt.Println("猫动")
}

func (c cat) eat(s string) {
	fmt.Printf("猫吃 %s \n", s)
}

func main() {

	var a1 animal // 定义一个接口类型的变量
	// a1
	// 动态类型	|-----------|
	// 			| nil		|
	// 			|-----------|
	// 动态值	|-----------|
	// 			| nil		|
	// 			|-----------|
	fmt.Printf("%T\n", a1) // <nil>

	bc := cat{
		name: "cat blue",
		feet: 0,
	}

	// 接口保存的时候分为两部分：值的类型和值本身
	// 接口类型： 动态类型和动态值
	a1 = bc
	a1.eat("鱼")
	fmt.Println(a1)
	fmt.Printf("%T\n", a1) // main.cat
}
