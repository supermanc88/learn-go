package main

// 反射 reflect
// 反射是指在程序运行期对程序本身进行访问和修改的能力。
// 程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。
// 在运行程序时，程序无法获取自身的信息

// 支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，
// 并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json :"age"`
}

type cat struct {
	name string
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type :%v\n", v)
	fmt.Printf("type name :%v type kind : %v\n", v.Name(), v.Kind())
	fmt.Printf("type value :%v\n", x)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()

	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	case reflect.Struct:
		fmt.Printf("type is struct, value is %v\n", v)
	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改会引发panic
	}
}

// 用反射传指针的时候，需要使用Elem这个函数
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	str := `{"name":"chm", "age":18}`
	var p person

	// 程序为什么知道person结构体中有 Name 和 Age这些字段？
	// 这就用到了反射
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)

	var a float64 = 3.14
	var b int64 = 100
	reflectType(a)
	reflectType(b)

	c := cat{"chm"}
	fmt.Println(c)
	reflectType(c)

	reflectValue(a)
	reflectValue(b)
	reflectValue(c)

	reflectSetValue2(&b)
	fmt.Println(b)
}
