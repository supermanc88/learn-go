package main

import "fmt"

// 结构体是值类型

type person struct {
	name string
	sex  string
}

func f1(x person) {
	x.sex = "woman"
}

func f2(x *person) {
	//(*x).sex = "woman"
	x.sex = "woman" // 和上面等价，语法糖，自动根据指针找对应的变量
}

func main() {
	var p1 person // 不是引用类型
	p1.name = "chm"
	p1.sex = "man"

	f1(p1)
	fmt.Println(p1) // {chm man}

	f2(&p1)
	fmt.Println(p1) // {chm woman}

	// 结构体指针1
	var p2 = new(person)
	p2.name = "xiaoxiao"
	fmt.Printf("%T\n", p2)  // *main.person
	fmt.Printf("%#v\n", p2) // &main.person{name:"xiaoxiao", sex:""}

	fmt.Printf("%p\n", p2)

	// 结构体指针2
	// 2.1 key - value 初始化
	var p3 = &person{
		name: "haha",
		sex:  "man",
	}

	fmt.Printf("%v\n", p3)

	// 2.2 使用值列表的形式，声明顺序要和定义时排字段的顺序一致
	p4 := &person{
		"aaa",
		"bb",
	}
	fmt.Printf("%v\n", p4)

	// 结构体的内存布局
	fmt.Printf("%p %p\n", &p4.name, &p4.sex) // 0xc000004580 0xc000004590
}
