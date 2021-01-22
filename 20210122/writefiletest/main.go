package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 写文件
// os.OpenFile 函数可以以指定模式打开文件，从而实现文件写入相关功能
// func OpenFile(name string, flag int, perm FileMode) (*File, error) {
// 		...
// }
// 	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
//	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
//	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
//	// The remaining values may be or'ed in to control behavior.
//	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
//	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
//	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
//	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
//	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.

// 第三个参数在windows下无效，随意即可，在linux下是文件的权限
func writeDemo1() {
	fileObj, err := os.OpenFile("E:/Go/src/github.com/supermanc88/learn-go/20210122/writefiletest/wtest.txt",
		os.O_CREATE | os.O_APPEND | os.O_TRUNC, 0644)

	if err != nil {
		fmt.Printf("OpenFile failed err:%v\n", err)
		return
	}

	defer fileObj.Close()

	fileObj.Write([]byte("mengbile\n"))

	fileObj.WriteString("直接写字符串")
}


func writeDemo2() {

	fileObj, err := os.OpenFile("E:/Go/src/github.com/supermanc88/learn-go/20210122/writefiletest/wtest.txt",
		os.O_CREATE | os.O_APPEND | os.O_TRUNC, 0644)

	if err != nil {
		fmt.Printf("OpenFile failed err:%v\n", err)
		return
	}

	defer fileObj.Close()

	// 创建一个写对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("bufio write string")

	// 上面是把内容写到了缓存中，需要再刷新一下
	wr.Flush()
}

func writeDemo3() {
	str := "hello ioutil writefile"

	err:= ioutil.WriteFile("E:/Go/src/github.com/supermanc88/learn-go/20210122/writefiletest/wtest.txt", []byte(str), 0666)

	if err != nil {
		fmt.Printf("ioutil writefile failed error :%v\n", err)
		return
	}
}

func main() {
	writeDemo3()
}
