package main

import "fmt"

// goto 语句			goto + label

func gotoDemo() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Printf("%v - %v\n", i, j)
		}
	}
	return
breakTag:
	fmt.Println("goto到这里")
}

// break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上
func breakDemo() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break	// 只这样的话，应该是跳出内层循环
				break BREAKDEMO1
			}
			fmt.Printf("%v - %v\n", i, j)
		}
	}
	fmt.Println("....")

}

// continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。
func continueDemo() {
forloop1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//continue	// 这种是忽略下面的代码，继续执行下次循环
				continue forloop1
			}
			fmt.Printf("%v - %v\n", i, j)
		}
	}
	fmt.Println("....")
}

func main() {
	gotoDemo()

	breakDemo()

	continueDemo()
}
