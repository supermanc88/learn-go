package main

import "fmt"

// 在go中只有一种循环，就是for循环

//  for 初始语句；条件表达式； 结束语句 {
//		循环体
// }

func main() {

	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种1
	//var i = 5
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}

	// 变种2 结束语句也可省略
	var i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 死循环
	//for {
	//	fmt.Println("aaaaaaa")
	//}

	// for range 遍历数组、切片、字符串、map及通道(channel)
	// 返回值有以下规律
	// 1. 数组、切片、字符串返回索引和值
	// 2. map返回键和值
	// 3. 通道只返回通道内的值

	s := "hello chm"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	for _, v := range s {
		fmt.Printf("%c\n", v)
	}

	// 跳出for循环 break
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		if j == 5 {
			break
		}
	}
	fmt.Println("over")

	// 当j = 5 时 跳过for循环
	for j := 0; j < 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println(j)
	}
	fmt.Println("continue over")
}
