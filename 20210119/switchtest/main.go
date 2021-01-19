package main

import "fmt"

// switch 语句
// 简化大量的判断

func main() {
	var n = 5

	// 不需要break
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效")
	}

	// 变种1
	switch m := 7; m {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println(m)
	}

	// 变种2
	age := 30
	switch {
	case age < 25:
		fmt.Println("小于25")
	case age > 25 && age < 35:
		fmt.Println("好好工作")
	case age > 60:
		fmt.Println("好好享受")
	default:
		fmt.Println("活着真好")
	}

	// 变种3，fallthrough 可以执行满足条件的case的下一个case，是为了兼容c语言中的case设计的，一般不使用
	s := 'a'
	switch {
	case s == 'a':
		fmt.Println("a")
		fallthrough // 向下穿透
	case s == 'b':
		fmt.Println("b")
	case s == 'c':
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
