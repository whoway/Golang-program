package main

import (
	"fmt"
)

func main() {
	// 第1种声明方式
	// 声明myMap1是一种map类型 key是string， value是string
	// key在左，value在右
	// var myMap map[key]value
	// 下面的仅仅是声明，里面没有任何的实例空间
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	// 在使用map前， 需要先用make给map分配数据空间，并且还能指定cap
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// 第2种声明方式，用make给map分配数据空间
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	// 【注意】下面的打印告诉我们，顺序不会被排序！因为Go的map只是基本的哈希
	fmt.Println(myMap2)

	// 第3种声明方式，
	myMap3 := map[string]string{
		"one":   "php",
		"two":   "c++",
		"three": "python",
	}
	fmt.Println(myMap3)
}
