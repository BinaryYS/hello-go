package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Println(unsafe.Sizeof(m))

	fmt.Printf("%#v\n", ptr)
	fmt.Println(unsafe.Sizeof(ptr))

	fmt.Printf("%#v\n", c)
	fmt.Println(unsafe.Sizeof(c))

	fmt.Printf("%#v\n", sl)
	fmt.Println(unsafe.Sizeof(sl))

	fmt.Printf("%#v\n", f)
	fmt.Println(unsafe.Sizeof(f))

	fmt.Printf("%#v\n", i)
	fmt.Println(unsafe.Sizeof(i))

}
