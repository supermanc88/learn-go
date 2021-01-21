package main

import (
	"fmt"
)

// 空接口
// interface{}

// 空接口没有必要起名字
// 所有的类型都实现了空接口
// 也是是能保存 任意类型 的变量
type xx interface {
}

func show(a interface{}) {

	// 类型断言
	// 用来判断空接口中传进来的是什么类型
	str, ok := a.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("not string")
	}

	fmt.Printf("type: %T, value: %v\n", a, a)
}

func show2(a interface{}) {

	// 类型断言
	switch t := a.(type) {
	case string:
		fmt.Printf("string: %s", t)
	case int:
		fmt.Println("int")
	default:
		fmt.Println("default")
	}
	fmt.Printf("type: %T, value: %v\n", a, a)
}

func main() {
	var m1 map[string]interface{}

	m1 = make(map[string]interface{}, 16)

	m1["name"] = "chm"
	m1["age"] = 18
	m1["haha"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}

	fmt.Println(m1)

	show(5)   // type: int, value: 5
	show(nil) // type: <nil>, value: <nil>
	show("123")

	show2(make(map[string]int, 10))
}
