package main

import "fmt"

// pointer
// 指针
// go中不存在指针操作，只需要两个符号：
// 1. & : 取地址
// 2. * : 根据地址取值

func main() {
	n := 18
	fmt.Println(&n)

	// 取地址
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)	// *int

	// 根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	// 为什么会panic
	// 和c中一样，指针没有赋值指向的内存地址
	//var a *int	// 相当于 a = nil
	//*a = 100
	//fmt.Println(*a)

	// new函数申请一个内存地址
	a := new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)
	fmt.Println(a)

	// make 也是用于分配内存的，区别于new，它只用于slice、map、chan的内存创建，而且它返回的类型就是这三个类型本身
	// 不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
	// new一般给基本的数据类型申请内存，int/string... 返回对应类型的指针 *string *int 很少用

}
