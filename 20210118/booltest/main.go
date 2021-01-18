package main

import "fmt"

// 布尔值
// 不允许将整型强制转换成bool类型
// bool类型无法参于数值运算，也无法与其他类型进行转换

func main() {
	var b1 bool				// bool类型默认为false
	b2 := true
	fmt.Println(b1, b2)

	fmt.Printf("%T value = %v", b1, b1)
}