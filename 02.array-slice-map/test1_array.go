package main

import (
	"fmt"
)

func printArray(myArray [4]int) {
	//值拷贝
	//【最坑的地方】go的数组也是值拷贝！！和cpp不一样！！
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}

	myArray[0] = 111
}

func main() {
	//固定长度的数组
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4} //这种后边没初始化的是0
	myArray3 := [4]int{11, 22, 33, 44}

	//下面是基本的for循环语法
	//for i := 0; i < 10; i++ {
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	//遍历数据的另一个方法，遍历这个还可以用range关键词
	// range的使用比较灵活，有的你如果遍历其他集合，会有不同返回
	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3)
	fmt.Println(" ------ ")
	for index, value := range myArray3 {
		fmt.Println("index = ", index, ", value = ", value)
	}
}
