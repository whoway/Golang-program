package main

import (
	"fmt"
)

func printArray(myArray []int) {
	//引用传递【最坑】Go的这个slice传入到函数里面是，引用传递！
	// _ 表示匿名的变量【学习这个！】
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	//【最坑】Go的这种不指定数组里面长度的方式，我们叫做slice，切片
	myArray := []int{1, 2, 3, 4} // 动态数组，切片 slice

	fmt.Printf("myArray type is %T\n", myArray)

	printArray(myArray)

	fmt.Println(" ==== ")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
