package first

import "fmt"

func main() {
	var a = 10

	switch a {
	case 5:
		fmt.Println("5")
	case 10:
		fmt.Println("10")			// 匹配到后面的就不再匹配，没有break
	case 15:
		fmt.Println("15")
	}

	fmt.Println("main...over....")
}
