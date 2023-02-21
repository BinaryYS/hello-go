package main

import (
	"fmt"
	"time"
)

func testGoroutine(a int) {
	count := 0
	for {
		count++
		fmt.Println(a)
	}
}
func main() {
	for i := 0; i < 100; i++ {
		go testGoroutine(i)
	}

	time.Sleep(1 * time.Second)
}
