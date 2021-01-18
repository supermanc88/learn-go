package main

import "fmt"

// 浮点数

func main() {
	//math.MaxFloat32			// float32 默认最大数
	f1 := 1.23456				// 默认类型是float64

	f2 := float32(1.23456)

	fmt.Printf("%T, %T\n", f1, f2)

	//f1 = f2					// float32不能赋值给float64
}