package main

import "fmt"

// defer
// go语言中的 defer 会将其后面跟随的语句进行延迟处理。
// 在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
// 也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行

// 一般用在资源释放的时候

// defer执行时机
// go中的return不是原子操作，可拆分
// 函数中return x 底层实现 --> 返回值 = x --> ret 指令
// defer执行时机 return x 底层实现 --> 返回值 = x --> 运行defer --> ret 指令

// result:
// start
// end
// heiheihei
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("heiheihei") // defer 把它后面的语句延迟到函数即将返回的时候再执行
	fmt.Println("end")
}

// result:
// end
// hahaha
// hehehe
// heiheihei
// start
func deferDemo1() {
	fmt.Println("start")
	defer fmt.Println("heiheihei") // 一个函数中可以有多个defer，执行顺序按照先进后出
	defer fmt.Println("hehehe")
	defer fmt.Println("hahaha")
	fmt.Println("end")
}

// 5 ret = x x++
func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x，不是返回值
	}()
	return x
}

// 6
// x = 5 x++
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// 5
// y = x = 5 x++
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 5
func f4() (x int) {
	defer func(x int) {
		x++ // 是操作的x的副本
	}(x)
	return 5
}

func main() {
	//deferDemo()
	//deferDemo1()

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
