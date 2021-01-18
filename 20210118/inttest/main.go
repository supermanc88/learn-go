package main

import "fmt"

// 整形

func main()  {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	// 转2进制输出
	fmt.Printf("%b\n", i1)
	// 转8进制输出
	fmt.Printf("%o\n", i1)
	// 转16进制输出
	fmt.Printf("%x\n", i1)

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 16进制
	i3 := 0x12345678
	fmt.Printf("%d\n", i3)


	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9)		// 明确指定类型，否则默认int类型
	fmt.Printf("%T\n", i4)

}