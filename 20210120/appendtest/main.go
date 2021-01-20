package main

import "fmt"

// append

func main() {
	a1 := [...]int{1,2,3,4,5,6,7,8}
	s1 := a1[:]

	fmt.Println(a1)
	fmt.Println(s1)
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(a1)
	fmt.Println(s1)

	// result:
	//	[1 2 3 4 5 6 7 8]
	//	[1 2 3 4 5 6 7 8]
	//	[1 3 4 5 6 7 8 8]
	//	[1 3 4 5 6 7 8]
}
