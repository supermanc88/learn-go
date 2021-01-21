package main

// 只有main包可以编译成一个可执行文件

// 包名是从 gopath/src 后面目录下开始的

import (
	"fmt"
	"github.com/supermanc88/learn-go/20210121/packagetest"
)

// 调用顺序：
// 全局声明 --> init --> main
//

func init() {
	fmt.Println("auto running...")
}

func main() {
	ret := packagetest.Add(10, 20)

	fmt.Println(ret)
}
