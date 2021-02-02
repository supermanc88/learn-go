package main

import "fmt"

func main() {
	// copy复制切片
	sss1 := []int{1, 2, 3}
	sss2 := sss1
	var sss3 = make([]int, 3, 5)
	fmt.Println(sss3) // [0 0 0]
	copy(sss3, sss1)
	fmt.Println(sss1, sss2, sss3) // [1 2 3] [1 2 3] [1 2 3]
	sss1[0] = 111
	fmt.Println(sss1, sss2, sss3) // [111 2 3] [111 2 3] [1 2 3]
}
