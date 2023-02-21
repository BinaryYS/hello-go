package main

import (
	"fileop"
	"fmt"
)

func main() {
	fileContent := fileop.Read()
	fmt.Printf("获取的内容是:%s\n", fileContent)

	//fileContent = fileop.ReadFile("/Users/ssg/Desktop/trade.jmx")
	//fmt.Printf("获取的内容2是:%s\n", fileContent)

	fmt.Println("hello, ssg, welcome to go world. best wished to you")
	var score int = 90
	sum, sub := sumCal(score, 110)
	fmt.Println(sum)
	fmt.Println(sub)

	sum1, _ := sumCal(score, 22)
	fmt.Println(sum1)
	fmt.Println(changeVal(1, 2))

	fmt.Println("可变参数")
	test(1, 10)
	test(1, 10, 12)

	fmt.Println("通过指针修改变量")
	a, b := 1, 5
	changeVal1(&a, &b)
	fmt.Println(a, b)

	fmt.Println("函数类型")
	t := changeVal
	fmt.Printf("t的类型%T, %T\n", t, changeVal)
	t(10, 21)

	fmt.Println("函数作为函数参数")
	funcPara(a, test1)

	fmt.Println("函数别名")

}

type myFunc func(val int)

func test1(num int) {
	fmt.Println("==================", num)
}
func funcPara(a int, testFunc func(int)) {
	fmt.Println("funcPara")
	test1(a)
}
func sumCal(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

/**
 *值传递
 */
func changeVal(a int, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

//可变参数
func test(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Print(args[i])
		fmt.Print(",")
	}
	fmt.Println("")
}

func changeVal1(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
