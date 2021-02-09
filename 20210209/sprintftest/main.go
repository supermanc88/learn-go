package main

import "fmt"

// sprintf

func main() {
	str1 := "10.20.88.223"
	str2 := "20001"

	str := fmt.Sprintf("add %s:%s", str1, str2)

	fmt.Println(str)

	pi := 3.14

	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", str1, pi, str2)
	fmt.Println(variant)

	// 匿名结构体声明，并赋予初值
	profile := &struct {
		name string
		hp   int
	}{
		name: "rat",
		hp:   150,
	}

	fmt.Printf("%+v\n", profile)
	fmt.Printf("%#v\n", profile)
	fmt.Printf("%T\n", profile)
}
