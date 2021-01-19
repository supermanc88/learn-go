package main

import (
	"fmt"
	"strings"
)

// string

func main() {

	path := "E:\\Go\\src\\github.com\\supermanc88\\learn-go\\20210119\\"

	s1 := "i'm ok"

	fmt.Println(path)
	fmt.Println(s1)

	// 多行字符串 其中内容原样输出
	s2 := `
哈哈只
	测试多行
`
	fmt.Println(s2)

	// 字符串相关操作 长度
	fmt.Println(len(s2))

	// 字符串拼接
	name := "chm"
	world := "dsb"
	ss := name + world
	fmt.Println(name + world)
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Printf("%s%s\n", name, world)
	fmt.Println(ss1)

	// 分割
	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	// 是否包含
	fmt.Println(strings.Contains(path, "github.com"))
	fmt.Println(strings.Contains(path, "test.com"))

	// 是否有前缀和后缀
	fmt.Println(strings.HasPrefix(path, "E:"))
	fmt.Println(strings.HasSuffix(path, "aaa"))

	// 子串位置
	s4 := "abcdef"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	// 拼接 切片
	fmt.Println(strings.Join(ret, "+"))

	// 字符串修改，
	// 字符串是不能修改的，需要转换成别的类型才修改
	ss2 := "白萝卜"
	ss3 := []rune(ss2) // 把字符串强制转换成了一个rune切片，切片中保存的就是字符,rune是int32类型的别名
	ss3[0] = '红'
	fmt.Println(ss2)
	fmt.Println(string(ss3)) //把rune切片强制转换成string类型

	// 字符吓字符串的区别
	c1 := "红"
	c2 := '红'
	fmt.Printf("%T %T\n", c1, c2)

	ss4 := "hello world"
	ss5 := []rune(ss4)
	fmt.Printf("%T %T\n", ss4, ss5[0])

	// 类型转换
	// 整型和float转、字符和字符串转
	n := 10 // int
	var f float64
	f = float64(n)
	fmt.Println(f)

	testStr := "hello沙河小王子"
	var hanNum int

	for _, v := range testStr {
		if v > 255 {
			hanNum++
		}
	}
	fmt.Printf("其中的汉字有 %d 个\n", hanNum)
}
