package main

import (
	"errors"
	"fmt"
)

/**
如何包装错误，defer作用范围， 最外层怎么捕获异常，每层都往外抛好麻烦
*/
func main() {
	//fmt.Println(testError(2, 0))
	result, err := testError2(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result=", result)
	}

	fmt.Println("异常被捕获，继续执行代码")
}

func testPanic() {

}

/**
defer + recover捕获处理异常机制
*/
func testError(a int, b int) int {
	defer func() {
		// 调用内置函数可以捕获异常
		err := recover()
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println(err)
		}
	}()

	result := a / b
	return result
}

/**
  自定义错误
*/
func testError2(a int, b int) (result int, err error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	} else {
		return a / b, nil
	}
}
