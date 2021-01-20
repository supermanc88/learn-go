package main

import "fmt"

// fmt

func main() {
	fmt.Println("beijing")
	fmt.Print("shanghai\n")

	name := "chm"
	fmt.Printf("%s\n", name)

	m1 := make(map[string]int, 1)
	m1["shang"] = 1
	fmt.Printf("%v\n", m1)  // map[shang:1]
	fmt.Printf("%#v\n", m1) // map[string]int{"shang":1}

	num := 90
	// %转义
	fmt.Printf("%d%%\n", num)

	b := true
	fmt.Printf("%t\n", b)

	fmt.Printf("%q\n", name) // "chm" 转成一个带”“的字符串

	// 获取用户输入
	var s string
	fmt.Scan(&s)
	fmt.Println("输入的内容是：", s)

	var (
		name1 string
		age   int
		class string
	)
	//fmt.Scanf("%s %d %s\n", &name1, &age, &class)
	fmt.Scanln(&name1, &age, &class)
	fmt.Println(name1, age, class)

}
