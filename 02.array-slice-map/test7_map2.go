package main

import (
	"fmt"
)

// 【注意】Go里面的map传入到函数是引用传递！
func printMap(cityMap map[string]string) {
	//cityMap 是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	// 1.添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 2.遍历
	printMap(cityMap)

	// 3.删除，使用delete函数，我们是传入key来删除整个key value的对的
	delete(cityMap, "China")

	// 4.修改
	cityMap["USA"] = "DC"
	ChangeValue(cityMap)

	fmt.Println("-------")

	// 遍历
	printMap(cityMap)
}
