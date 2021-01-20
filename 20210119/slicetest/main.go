package main

import (
	"fmt"
)

// slice 切片
// 数组的长度是固定的，有局限性
// 切片是一个拥有相同类型元素的可变长度的序列。
// 它是基于数组类型做的一层封装，支持自动扩容
// 切片是一个引用类型，它的内部结构包含  地址、长度和容量
// 一般用于快速地操作一块数据集合
// 语法：
// var name []T

// 总结：
// 切片指向了一个底层数组
// 切片的长度是它的元素个数
// 切片的容量是底层数组从切片的第一个元素到底层数组最后一个的数量
// 切片就是一个框，框住了一块连续的内存
// 真正的数据都是保存在底层数组中的

func main() {

	var s1 []int //定义一个存放int类型的切片
	var s2 []string

	fmt.Println(s1, s2)

	fmt.Println(s1 == nil) // nil就是空，就是没有开辟内存空间
	fmt.Println(s2 == nil)

	// 初始化 和数组类似
	s1 = []int{1, 2, 3}
	s2 = []string{"12", "333", "444"}
	fmt.Println(s1, s2)

	// 切片的长度和容量，可以通过len求长度，也可以使用cap函数求切片的容量
	fmt.Printf("len: %d, cap: %d\n", len(s1), cap(s2))

	// 2.由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 右不包含，左闭右开
	fmt.Println(s3)
	fmt.Printf("len : %d, cap : %d\n", len(s3), cap(s3)) // len : 4, cap : 7

	s4 := a1[1:] // 从第1个切到最后 [1:len(a1)] [3,5,7,9,11,13]
	s5 := a1[:4] // 从第0个切到第4个 [0:4]
	s6 := a1[:]  // 全部 [0:len(a1)]

	fmt.Println(s4, s5, s6)

	// 切片的容量是指底层数组从切片的第一个元素到最后的数量
	fmt.Printf("s5 len:%d s5 cap:%d\n", len(s5), cap(s5)) // s5 len:4 s5 cap:7
	fmt.Printf("s4 len:%d s4 cap:%d\n", len(s4), cap(s4)) // s4 len:6 s4 cap:6

	// 3.切片再切片
	s7 := s4[3:]
	fmt.Println(s7)
	fmt.Printf("s7 len:%d s7 cap:%d\n", len(s7), cap(s7))

	// 修改底层数组元素的值
	fmt.Println(s4, s6, s7)
	a1[6] = 1300
	fmt.Println(s4, s6, s7)

	// 使用make函数构造切片
	// make的参数 []T 长度 容量
	ms1 := make([]int, 5, 10)
	fmt.Printf("%d %d %v\n", len(ms1), cap(ms1), ms1)

	ms2 := make([]int, 0, 10)
	fmt.Printf("%d %d %v\n", len(ms2), cap(ms2), ms2)

	// 切片不能直接比较，不能使用 == 操作来判断两个切片是否含有全部相等的元素。
	// 切片唯一合法的比较操作是和 nil 比较
	// 一个 nil 值的切片并没有底层数组，一个 nil 值的切片长度和容量都是0
	// 但不能说一个长度和容量都是0的切片一定是 nil
	// 所以要判断一个切片是否为空应该用len判断
	var nilt1 []int         // len 0 cap 0 s1 == nil
	nilt2 := []int{}        // len 0 cap 0 s1 != nil
	nilt3 := make([]int, 0) // len 0 cap 0 s1 != nil
	fmt.Println(nilt1 == nil)
	fmt.Println(nilt2 == nil)
	fmt.Println(nilt3 == nil)

	// 切片的赋值拷贝
	copyt1 := []int{1, 3, 5}
	copyt2 := copyt1 // 两个都指向了同一个底层数组
	fmt.Println(copyt1, copyt2)
	copyt1[0] = 2
	fmt.Println(copyt1, copyt2)

	// 切片的遍历和数组的遍历一样
	for i := 0; i < len(copyt1); i++ {
		fmt.Println(copyt1[i])
	}

	for i, v := range copyt1 {
		fmt.Println(i, v)
	}

	// append 为切片追加元素
	// append()
	appendtest := []string{"beijing", "shanghai", "shenzhen"}
	//appendtest[3] = "guangzhou"	不能这样添加 index out of range
	fmt.Printf("%v %d %d\n", appendtest, len(appendtest), cap(appendtest)) // [beijing shanghai shenzhen] 3 3

	// 调用append函数必须用原来的切片变量接收返回值
	// 原因是底层数组保存不了了之后，会把底层数组换掉，所以切片需要指向新的地址
	// 扩容策略
	appendtest = append(appendtest, "guangzhou")
	fmt.Printf("%v %d %d\n", appendtest, len(appendtest), cap(appendtest)) // [beijing shanghai shenzhen guangzhou] 4 6

	appendtest = append(appendtest, "hangzhou", "chengdu")
	fmt.Printf("%v %d %d\n", appendtest, len(appendtest), cap(appendtest))

	sss := []string{"wuhan", "xian", "suzhou"}
	appendtest = append(appendtest, sss...) // ...表示拆开
	fmt.Printf("%v %d %d\n", appendtest, len(appendtest), cap(appendtest))

	// copy复制切片
	sss1 := []int{1, 2, 3}
	sss2 := sss1
	var sss3 = make([]int, 3, 5)
	copy(sss3, sss1)
	fmt.Println(sss1, sss2, sss3) // [1 2 3] [1 2 3] [1 2 3]
	sss1[0] = 111
	fmt.Println(sss1, sss2, sss3) // [111 2 3] [111 2 3] [1 2 3]

	// 从切片中删除元素
	// go中没有删除切片元素的专用方法
	rms1 := []int{1, 2, 3, 4, 5, 6, 7}
	// 删除索引为2的元素 其实就是拆成两个切片再拼接
	// 删除索引为index的元素 a = append(a[:index], a[index+1:]...)
	fmt.Println("before remove cap:", cap(rms1))
	// 删除内容会影响底层数组存储的内容
	rms1 = append(rms1[:2], rms1[3:]...)
	fmt.Println(rms1)
	fmt.Println("after remove cap:", cap(rms1))

	aa1 := [...]int{1, 2, 3, 4, 5} // 数组
	qq1 := aa1[:]                  // 切片
	fmt.Printf("qq1 addr %p\n", &qq1[0])
	qq1 = append(qq1[:1], qq1[2:]...)
	fmt.Printf("qq1 addr %p\n", &qq1[0])
	fmt.Printf("after remove cap: %d %v\n", cap(qq1), qq1) // after remove cap: 5 [1 3 4 5]
	fmt.Println(aa1)                                       // [1 3 4 5 5]	修改了底层数组

	var at = make([]int, 5, 10) // 注意这里的长度是5，就说明已经有了一个切片[0,0,0,0,0] 之后在后面追加
	for i := 0; i < 10; i++ {
		at = append(at, i)
	}
	fmt.Println(at) // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

}
