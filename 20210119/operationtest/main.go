package main

import "fmt"

// 运算符

func main() {
	var (
		a = 5
		b = 2
	)

	// 算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// 这是单独的语句，不能放在=号的右边赋值   a = a++ 这样是不可以的
	a++
	b--

	// 关系运算符 返回true或false
	// ==
	// !=
	// >
	// <
	// <=
	// >=

	fmt.Println(a == b) // go是强类型的，相同类型的变量才能比较 如：float32和float64是不能比较的

	// 逻辑运算符	返回true或false
	// &&
	// ||
	// !

	age := 22

	if age > 18 && age < 60 {
		fmt.Println("上班的")
	} else {
		fmt.Println("不用上班的")
	}

	if age < 18 || age > 60 {
		fmt.Println("不用上班的")
	} else {
		fmt.Println("上班的")
	}

	isMarried := false

	if !isMarried {
		fmt.Println("没有结婚")
	}

	// 位运算符
	// &		两位均为1才为1
	// |		两位有一个为1就为1
	// ^		两位不同则为1
	// <<		a << b 把a的2进制向左移动b位，高位丢弃，低位补0
	// >>		a >> b 把a的2进制向右移动b位

}
