package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 必须加随机种子，不然每次执行生成的随机数都一样
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(100)
		fmt.Println(r1, r2)
	}
}
