package main

import "fmt"

// 数组的初始化

func main() {
	var a1 [10]int
	fmt.Println(a1)

	var a2 = [5]int{1, 2}
	fmt.Println(a2)

	// 根据初始化的长度自动确认数组长度
	var a3 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a3)
}
