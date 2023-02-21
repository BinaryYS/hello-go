package main

import (
	"container/list"
	"fmt"
)

func main() {

	list := list.New()
	list.PushBack("ssg")
	list.PushBack("zmn")
	printList(list)

	fmt.Println(list.Len())

	// 尾部添加
	list.PushBack("canon")
	printList(list)
	// 头部添加
	list.PushFront("header")
	printList(list)
	// 尾部添加后保存元素句柄
	element := list.PushBack("first")
	printList(list)
	// 在fist之后添加high
	list.InsertAfter("high", element)
	printList(list)
	// 在fist之前添加noon
	list.InsertBefore("noon", element)
	printList(list)
	// 使用
	list.Remove(element)
	printList(list)
}

func printList(list *list.List) {
	for element := list.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value, "->")
	}
	fmt.Println("")
}
