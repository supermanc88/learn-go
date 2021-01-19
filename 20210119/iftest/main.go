package main

import "fmt"

// if 条件判断

func main() {
	//age := 19
	//
	//if age > 18 {
	//	fmt.Println("澳门首家线上赌场开业了")
	//} else {
	//	fmt.Println("该写暑假作业了")
	//}
	//
	//// 多个条件判断
	//if age > 35 {
	//	fmt.Println("人到中年")
	//} else if age > 18 {
	//	fmt.Println("青年")
	//} else {
	//	fmt.Println("好好学习")
	//}

	// 作用域
	// 此时age只在if条件判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("澳门首家线上赌场开业了")
	} else {
		fmt.Println("该写暑假作业了")
	}
}
