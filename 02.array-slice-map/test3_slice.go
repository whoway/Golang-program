package main

import (
	"fmt"
)

func main() {
	//声明方式1：声明slice1是一个切片，并且初始化，默认值是1，2，3。 长度len是3
	//slice1 := []int{1, 2, 3}

	//声明方式2：声明slice1是一个切片，但是【并没有给slice分配空间】
	var slice1 []int
	//slice1 = make([]int, 3) //开辟3个空间 ，默认值是0

	//声明方式3：声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值是0
	//var slice1 []int = make([]int, 3)

	//声明方式4：声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值是0, 通过:=推导出slice是一个切片
	//slice1 := make([]int, 3)

	// %v表示打印一些详细信息，比如{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	//判断一个silce【是否为空，啥元素都没有】
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
