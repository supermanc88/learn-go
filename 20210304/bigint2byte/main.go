package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigdata := new(big.Int)

	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	bigdata.SetBytes(data)

	// 5753854965885600108563692893507482093143722240
	// [1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0]
	fmt.Println(bigdata)

	fmt.Println(bigdata.Bytes())
}
