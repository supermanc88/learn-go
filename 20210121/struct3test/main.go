package main

import "fmt"

// 结构体模拟实现继承

type animal struct {
	name string
}

type dog struct {
	feet   uint8
	animal // animal拥有的方法 dog也有
}

func (a animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}

func (d dog) wang() {
	fmt.Printf("%s 在叫：wangwang\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "chm",
		},
	}
	fmt.Println(d1)

	d1.wang()

	d1.move()
}
