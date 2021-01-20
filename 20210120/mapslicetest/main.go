package main

import "fmt"

// map 和 slice组合
// 注意用的时候要初始化

func main() {
	// 元素类型为map的切片
	s1 := make([]map[int]string, 5, 10)

	// 在使用map时注意里面map的初始化
	s1[0] = make(map[int]string, 1)
	s1[0][100] = "a"

	fmt.Println(s1)

	// 值为切片类型的map
	m1 := make(map[string][]int, 10)
	m1["beijing"] = []int{1, 2, 3}
	m1["shanghai"] = make([]int, 1, 5)
	fmt.Println(m1)

}
