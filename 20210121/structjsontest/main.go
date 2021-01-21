package main

import (
	"encoding/json"
	"fmt"
)

// 结构体转json
// 1. 把go中的struct转成json
// 2. json转成go中可识别的struct类型

type person struct {
	Name string `json:"name" db:"name" ini:"name"` // 当json解析的时候，使用name去表示这个字段
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "chm",
		Age:  111,
	}

	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", string(b))

	// 反序列化
	str := `{"name":"chm", "age":11}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)

	fmt.Println(p2)
}
