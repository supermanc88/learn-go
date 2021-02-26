package main

// 读写文件测试
// https://blog.csdn.net/whatday/article/details/103938124

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readfile1() {
	data, err := ioutil.ReadFile("E:\\Go\\src\\github.com\\supermanc88\\learn-go\\20210223\\rwfiletest\\data.txt")
	if err != nil {
		fmt.Printf("ioutil ReadFile err = %s\n", err)
		return
	}
	fmt.Println(string(data))
}

func writefile1() {
	err := ioutil.WriteFile("E:\\Go\\src\\github.com\\supermanc88\\learn-go\\20210223\\rwfiletest\\data.txt", []byte("teststring"), 0666)
	if err != nil {
		fmt.Printf("ioutil WriteFile err %s\n", err)
		return
	}
	fmt.Println("ioutil WriteFile success")
}

func readfile2() {
	f, err := os.Open("E:\\Go\\src\\github.com\\supermanc88\\learn-go\\20210223\\rwfiletest\\data.txt")
	if err != nil {
		fmt.Printf("os Open err = %s\n", err)
		return
	}

	defer f.Close()

	var data []byte
	buf := make([]byte, 1024)

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("read file buf err = %s\n", err)
			return
		}

		if n == 0 {
			break
		}

		data = append(data, buf[:n]...)
	}

	fmt.Println(string(data))
}

func writefile2() {
	f, err := os.Open("E:\\Go\\src\\github.com\\supermanc88\\learn-go\\20210223\\rwfiletest\\data.txt")
	if err != nil {
		fmt.Printf("os Open err = %s\n", err)
		return
	}

	defer f.Close()

	n, err := f.Write([]byte("teststring"))
	fmt.Printf("write %d bytes\n", n)

	n, err = f.WriteString("writestring")
	fmt.Printf("write string %d bytes\n", n)

	f.Sync()
}

func readfile3() {
	f, err := os.Open("E:\\Go\\src\\github.com\\supermanc88\\learn-go\\20210223\\rwfiletest\\data.txt")
	if err != nil {
		fmt.Printf("os Open err = %s\n", err)
		return
	}

	defer f.Close()

	r := bufio.NewReader(f)
	var data []byte
	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return
		}
		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)
	}
	fmt.Println(string(data))
}

func writefile3() {
	f, err := os.Open("E:\\Go\\src\\github.com\\supermanc88\\learn-go\\20210223\\rwfiletest\\data.txt")
	if err != nil {
		fmt.Printf("os Open err = %s\n", err)
		return
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	n, err := w.WriteString("bufio new writer")
	fmt.Printf("bufio NewWriter write %d bytes\n", n)
	w.Flush()
}

func readfile4() {
	f, err := os.Open("E:\\Go\\src\\github.com\\supermanc88\\learn-go\\20210223\\rwfiletest\\data.txt")
	if err != nil {
		fmt.Printf("os Open err = %s\n", err)
		return
	}

	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("ioutil ReadAll err %s\n", err)
		return
	}

	fmt.Println(string(fd))
}

func writefile4() {
	f, err := os.Open("E:\\Go\\src\\github.com\\supermanc88\\learn-go\\20210223\\rwfiletest\\data.txt")
	if err != nil {
		fmt.Printf("os Open err = %s\n", err)
		return
	}

	defer f.Close()

	n, err := io.WriteString(f, "teststring")
	fmt.Printf("io WriteString write %d bytes\n", n)
}

func main() {
	readfile4()
}
