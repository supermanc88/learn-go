package main

import "fmt"

// map
// 映射关系容器 map，其内部使用散列表hash实现
// 无序的 基于 key value的数据结构
// map是引用类型，必须初始化之后才能使用

// 定义语法：
// map[keyType]valueType

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) // true 这是还没有开辟内存

	m1 = make(map[string]int, 10) // 要估算好map的容量，避免在程序运行期间再动态扩容

	m1["aaa"] = 111

	m1["bbb"] = 2222

	fmt.Println(m1)

	fmt.Println(m1["aaa"])

	fmt.Println(m1["ccc"])	// 0 不存在的key，拿到对应类型的零值

	// 检查是否存在此值
	value, ok := m1["ccc"]
	if !ok {
		fmt.Println("no this key")
	} else {
		fmt.Println(value)
	}


	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}


	// 删除键值对
	delete(m1, "aaa")

	// 删除不存在的
	delete(m1, "ccc")
	fmt.Println(m1)


	m1["ddd"] = 444
	m1["eee"] = 555
	m1["fff"] = 666

}
