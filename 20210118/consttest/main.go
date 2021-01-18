package main

import "fmt"

// 常量定义之后就不能修改了
// 在程序运行期间不可改变
const pi = 3.1415926

// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 如果批量声明，如果某一行没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)


// iota 常量计数器
// iota在const出现时会初始化为0
// const 中每新加一行，iota就加一次
const (
	a1 = iota		// 0
	a2				// 1
	a3				// 2
)

const (
	b1 = iota		// 0
	b2				// 1
	_				// 2
	b3				// 3
)

// 插队
const (
	c1 = iota		// 0
	c2 = 100		// 100
	c3 = iota		// 2
	c4				// 100
)

// 多个常量声明在一行
const (
	d1, d2 = iota+1, iota+2		// d1 = 1 d2 = 2
	d3, d4 = iota+1, iota+2		// d3 = 2 d4 = 3
)

// 定义数量级
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
}
