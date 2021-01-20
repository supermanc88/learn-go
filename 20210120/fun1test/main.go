package main

import "fmt"

func fun1() {
	fmt.Println("fun1")
}

func fun2(name string) {
	fmt.Println("hello", name)
}

func fun3(x, y int) int {
	return x + y
}

func fun4(x, y int) (int, int) {
	return x + y, x - y
}

func main() {

}
