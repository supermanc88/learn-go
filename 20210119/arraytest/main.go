package main

import (
	"fmt"
)

// 数组 存放元素的容器
// var a [5]int
// 数组的长度必须是常量，长度是数组类型的一部分。
// 数组的长度不可改变
// [5]int 和 [10]int 是不同的类型
// 数组的类型和容量是这个类型的一部分

func main() {
	var a1 [3]bool
	var a2 [4]bool

	//a1 = a2		这是两个完全不同的类型

	fmt.Printf("%T %T\n", a1, a2) // [3]bool [4]bool

	// 数组的初始化
	// 如果不初始化，默认元素都是零值（bool是false，整型和浮点是0，字符串是“”）
	fmt.Println(a1, a2)

	// 1.初始方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	// 2. 初始方式2
	// 根据初始值自动推断数组长度
	//a10 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)

	// 3.方式3
	// 后面默认零值
	a3 := [5]int{1, 2}
	fmt.Println(a3)

	// 4.方式4
	// 根据索引初始化，其余位置补零值
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)

	// 数组的遍历
	str1 := [...]string{"北京", "上海", "深圳"}
	city := [...]string{"测试", "aaa", "bbb"}

	fmt.Println(city)

	// 方法1
	for i := 0; i < len(str1); i++ {
		fmt.Println(str1[i])
	}

	// 方法2
	for index, value := range str1 {
		fmt.Println(index, value)
	}

	// 多维数组
	// [[1 2] [3 4] [5 6]]
	var aa1 [3][2]int // 首先aa1的长度是3 里面的类型是[2]int的数组 这样理解比较好
	// 初始化
	aa1 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}

	fmt.Println(aa1)

	// 多维数组遍历
	for _, v1 := range aa1 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
		fmt.Println()
	}

	// 数组是值类型，区别与引用类型
	// 数组支持 == ！= 操作符，因为内存问题被初始化过的
	// [n]*T表示指针数组，*[n]T表示数组指针
	// 指针数组是数组中存放的指针
	// 数组指针代表的是这个数组的地址
	b1 := [3]int{1, 2, 3}
	b2 := b1 // b2是新的空间，不是指向于b1
	b2[0] = 100
	fmt.Println(b1, b2)

	// 测试题，求数组元素的和
	arr1 := [...]int{1, 3, 5, 7, 8}
	arrSum := 0
	for _, v := range arr1 {
		arrSum += v
	}
	fmt.Println(arrSum)

	// 测试题，找出上面数组中和为8的两个元素的下标，如 0，3、1，2
	for i, v1 := range arr1 {
		for j, v2 := range arr1 {
			if v1+v2 == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
	fmt.Println()

	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[i]+arr1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
