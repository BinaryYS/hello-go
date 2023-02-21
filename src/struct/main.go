package main

import "fmt"

//定义老师结构体
type Teacher struct {
	Name   string
	Age    int
	School string
}

type Tec Teacher //取别名
type integer int

// 绑定方法
func (val *integer) changeVal() {
	*val = 250
}

//  结构体绑定方法
func (a Teacher) learnSchool() {
	fmt.Println(a.Name + " studied in " + a.School)
}

/**
底层编译器会自动加上*，
*/
func (a *Teacher) changeName() {
	a.Name = "张小白"
}
func main() {
	var t1 Teacher //
	fmt.Println(t1)
	t1.Name = "杨松"
	t1.Age = 29
	t1.School = "重庆大学"
	fmt.Println(t1)
	fmt.Println(t1.School)
	fmt.Println(t1.Age + 1)

	// 方法2
	var t = Teacher{"张美宁", 28, "西安邮电大学"}
	fmt.Println(t)

	var t2 = new(Teacher)
	t2.Name = "李四"
	t2.Age = 12
	t2.School = "清华大学"
	fmt.Println(*t2)
	// 方法传参，值传递
	t2.learnSchool()
	t2.changeName()
	fmt.Println(*t2)

	var i integer = 320
	fmt.Println("before val=", i)
	i.changeVal()
	fmt.Println("after val=", i)
}
