package main

import (
	"fmt"
)

// 接口
// 接口是一种类型,特殊的类型，规定了变量有哪些方法
// type 接口名 interface {
//	方法名1(参数1， 参数2...)(返回值1，返回值2...)
//	方法名2(参数1， 参数2...)(返回值1，返回值2...)
// 	...
// }

// 用来给 变量、参数、返回值等 设置类型
// 一个变量如果实现了接口中规定的所有方法，那么这个变量就实现了这个接口，可以称为这个接口的变量

type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型
}

// 在编程中会遇到以下场景：
// 不关心一个变量是什么类型，只关心能调用它的什么方法

type cat struct {
}

type dog struct {
}

type bmwcar struct {
}

type audicar struct {
}

type mover interface {
	move()
}

func (b bmwcar) move() {
	fmt.Println("bmw move")
}

func (a audicar) move() {
	fmt.Println("audi move")
}

func drive(x mover) {
	x.move()
}

func (c cat) speak() {
	fmt.Println("miao")
}

func (c cat) move() {
	fmt.Println("cat move")
}

func (d dog) speak() {
	fmt.Println("wang")
}

// 约束传进来变量是否有指定的方法
func hit(x speaker) {
	// 接收一个参数，传进来谁打谁
	x.speak()
}

func main() {
	c1 := cat{}
	d1 := dog{}

	hit(c1)
	hit(d1)

	b1 := bmwcar{}
	a1 := audicar{}

	drive(b1)
	drive(a1)

	var ss speaker
	ss = c1
	fmt.Printf("%T, %v\n", ss, ss) // main.cat, {}
}
