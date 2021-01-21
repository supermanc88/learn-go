package packagetest

import "fmt"

// 包中的标识符(变量名、结构体、接口等)
// 如果首字母是小写的，表示私有 只能在当前这个包中使用
// 首字母大写可以对别的包可见
func Add(x, y int) int {
	return x + y
}

func init() {
	fmt.Println("import 我时 自动加载... ")
}
