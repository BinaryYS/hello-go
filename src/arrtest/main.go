package main

import (
	"fmt"
	"strings"
)

func main() {
	// 切片
	testStringsJoin()

	// 数组初始化
	//var scores [5]int = [5]int{3, 4, 9, 2, 2}
	//var scores1 = [5]int{8, 3, 1, 11, 11}
	//var arr3 = [...]int{1, 2, 3}
	//fmt.Println(arr3)
	//changeVal(&arr3)
	//fmt.Println(arr3)
	//var arr4 = [...]int{0: 21, 1: 99, 3}
	//fmt.Println(arr4)
	//for i := 0; i < 5; i++ {
	//	fmt.Println("请录入第%d个学生的成绩", i+1)
	//	fmt.Scanln(&scores[i])
	//}
	//
	//for i := 0; i < len(scores); i++ {
	//	fmt.Println("成绩为：", scores[i])
	//}

	//for key, val := range scores {
	//	fmt.Printf("第%d个学生成绩成绩为：%d\n", key, val)
	//}
	//
	//for key, val := range scores1 {
	//	fmt.Printf("第%d个学生成绩成绩为：%d\n", key, val)
	//}
	////arrSum()
	//twoDimensionalArr()
}

func testStringsJoin() {
	var strArr []string = []string{"zhangsan", "lisi", "wangwu"}
	fmt.Println(strArr)

	fmt.Println(strings.Join(strArr[0:], ","))
}
func testArr() {
	var intarr []int = []int{1, 4, 2, 4, 5, 7}
	var slice []int = intarr[1:3]
	slice2 := intarr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice)) // 切片元素个数
	fmt.Println(cap(slice)) // 切片容量 动态变化

	fmt.Println(slice2) //

	// 追加是底层扩容数组
	fmt.Println(intarr)
	slice3 := append(slice2, 88, 99)
	fmt.Println(slice3)
	fmt.Println(intarr)

	// copy
	var b []int = make([]int, 10)
	copy(b, intarr)
	fmt.Println(b)
}

func twoDimensionalArr() {
	var arr [2][3]int16
	fmt.Println(arr)

	fmt.Printf("arr起始地址%p\n", &arr)
	fmt.Printf("arr[0]起始地址%p\n", &arr[0])
	fmt.Printf("arr[0][0]起始地址%p\n", &arr[0][0])

	fmt.Printf("arr[1]起始地址%p\n", &arr[1])
	fmt.Printf("arr[1][0]起始地址%p\n", &arr[1][0])

	//遍历1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}

	fmt.Println("----------------")
	//遍历2
	for _, val := range arr {
		for _, v := range val {
			fmt.Print(v, "\t")
		}
		fmt.Println()
	}
}

func changeVal(arr *[3]int) {
	(*arr)[1] = -1
}
func arrSum() {
	var scores [4]float32
	scores[0] = 1
	scores[1] = 2
	scores[2] = 3
	scores[3] = 5

	var sum float32
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	var avg float32
	avg = sum / float32(len(scores))
	fmt.Printf("result sum=%v, avg=%v", sum, avg)

}
