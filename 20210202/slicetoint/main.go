package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// slice 转 int

func sliceToInt() {
	var byteArray = [2]byte{0x01, 0x02}

	bytesBuffer := bytes.NewBuffer(byteArray[:])

	var tmp uint16
	binary.Read(bytesBuffer, binary.LittleEndian, &tmp)

	fmt.Printf("0x%04x\n", tmp)
}

// int 转 slice
func intToSlice() {
	tmp := uint16(0x0201)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, &tmp)

	fmt.Println(bytesBuffer.Bytes())
}

func main() {
	sliceToInt()

	intToSlice()
}
