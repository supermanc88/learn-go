package first

import "fmt"

func main() {
	// 格式化打印
	// %v 原样输出
	// %T 打印类型

	a := 100
	b := 3.13
	c := true
	d := "hello world"
	e := `Runby`
	f := 'A'

	fmt.Printf("%T, %b\n", a, a)
	fmt.Printf("%T, %f\n", b, b)
	fmt.Printf("%T, %t\n", c, c)
	fmt.Printf("%T, %s\n", d, d)
	fmt.Printf("%T, %s\n", e, e)
	fmt.Printf("%T, %c\n", f, f)

	fmt.Printf("%v\n", f)


	var x int
	var y float64
	fmt.Scanln(&x, &y)
	fmt.Println(x, y)

}
