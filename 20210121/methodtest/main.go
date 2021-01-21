package main

import "fmt"

// 方法 method
// 是一种作用于特定类型变量的函数
// 这种特定类型变量叫做接收者 receiver
// 接收者的概念就类似于其他语言中的this或者self

// func (接收者变量 接收者类型) 方法名 (参数列表) (返回参数) {
// 		函数体
// }

// go中如果标识符首字母大写，就表示对外部可见(暴露的，公有的)

type dog struct {
	name string
	age  int
}

func newDog(name string) dog {
	return dog{name: name}
}

// 方法是作用于特定类型的函数
// 在函数面前指定一个谁可以使用这个函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s : wangwangwang\n", d.name)
}

// 需要修改接收者中的值
// 接收者是拷贝代价比较大的大对象
// 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者
func (d dog) guonian() {
	d.age++
}

func (d *dog) guonian2() {
	d.age++
}

func main() {
	d1 := newDog("chm")

	d1.wang()
	fmt.Println(d1.age)

	d1.guonian()
	fmt.Println(d1.age)

	d1.guonian2()
	fmt.Println(d1.age)
}
