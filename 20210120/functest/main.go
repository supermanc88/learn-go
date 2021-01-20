package main

import (
	"fmt"
)

// function 函数
// 格式：
// func 函数名 [(变量)] [(返回值)] {
// 		函数体
// }

// 函数1
func sum(x int, y int) (ret int) {
	return x + y
}

// 函数2
// 无返回值
func sum1(x int, y int) {
	fmt.Println(x + y)
}

// 函数3
// 无参数无返回值
func sum2() {
	fmt.Println("sum2")
}

// 函数4
// 无参数有返回值
func sum3() (ret int) {
	ret = 3
	return
}

// 参数可命名也可不命名
// 命名的优势在于可以直接在函数体中使用，如函数3中的使用
func sum4() int {
	//return 4
	ret := 4
	return ret
}

// 多个返回值
func sum5() (int, string) {
	return 1, "beijing"
}

// 参数的类型简写
// 当参数中连续多个参数类型一致时，可以将非最后一个类型省略
func sum6(x, y int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数
// y可不传、可多个，必须放在最后一个
func sum7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y的类型是个slice 切片
}

// go中没有默认参数这个概念

func main() {

	s := sum(1, 2)
	fmt.Println(s)
	sum7("beijing", 1, 2, 3, 4, 56)
}
