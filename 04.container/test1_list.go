package main

// 双向链表的使用
// Go的双向链表(实际是一个环)
import (
	"container/list"
	"fmt"
)

func main() {

	// 使用list.List{}延迟初始化
	l2 := list.List{}
	l2.PushFront(2)
	fmt.Println(l2.Front().Value) // 2
	fmt.Println("---------------")

	// 使用api+使用list.New()直接初始化
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		fmt.Println(iter.Value)
	}
	fmt.Println("---------------")
	// 头部添加
	l.PushFront(67)

	for iter := l.Front(); iter != nil; iter = iter.Next() {
		fmt.Println(iter.Value)
	}
	fmt.Println("---------------")

	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)

	for iter := l.Front(); iter != nil; iter = iter.Next() {
		fmt.Println(iter.Value)
	}
	fmt.Println("---------------")
}
