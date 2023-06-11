package main

import (
	"fmt"
)

func main() {

	s := []int{11, 22, 33} //len = 3, cap = 3, [1,2,3]

	// 注意：截取的语法是左闭右开的[0, 2)，所以下面是选择的是[11, 22]
	// 特殊的，s1 := [:]这种就是获取全部的元素
	// 特殊的，s1 := [1:]显然是第2个元素到最后的元素
	// 特殊的, s2 := [:4]显然是最开始的元素到下标为4-1位置的元素
	s1 := s[0:2]
	fmt.Println(s1)

	// 【最坑】：实际上s和s1是指向同一个底层数组！
	// 所以下面的改写，s和s1都会改变
	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	// 如果我们想要数组之间分开去指向，我们可以先新建slice，然后使用copy
	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) //s2 = [0,0,0]

	// copy(目标数组, 源数组)
	//将s中的值 依次拷贝到s2中
	copy(s2, s)
	s1[0] = 111
	s2[0] = 222
	fmt.Println(s)
	fmt.Println(s2)

}
