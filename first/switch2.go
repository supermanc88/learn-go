package main

import (
	"fmt"
)

func main() {
	score := 80
	// 无变量，默认为true
	switch {
	case score >= 0 && score < 60:
		fmt.Println("不及格")
	case score >= 60:
		fmt.Println("及格")
	}

	fmt.Println("-----------------------")

	letter := "A"

	switch letter {
	// case 可多少匹配项写一起
	case "A", "E", "I", "O", "U":
		fmt.Println("元音")
	case "H", "M":
		fmt.Println("H或M")
	default:
		fmt.Println("其他")
	}

	fmt.Println("--------------------------")
	// 变量定义在switch中
	switch lang := "go"; lang {
	case "go":
		fmt.Println("golang")
	case "python":
		fmt.Println("python")
	}
}
