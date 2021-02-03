package main

import (
	"encoding/json"
	"fmt"
)

// 自定义error

// type error interface {
//  Error() string
// }
// error 在标准库中被定义为一个接口类型，该接口只有一个Error方法
// 那么自定义error只要拥有Error方法，就实现了error接口

type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func main() {

	e := &Err{
		Code: 401,
		Msg:  "无权限",
	}
	fmt.Printf("%T\n", e)
	fmt.Println(e)
}
