package lib2

import (
	"fmt"
)

// 当前lib2包提供的API【如果1个函数的首字母大写字母，代表当前函数是个对外开发的函数】
func Lib2Test() {
	fmt.Println("lib2Test()...")
}

// 【如果1个函数的首字母小写字母，代表当前函数只能你当前包内调用！其他包调用不了！
func init() {
	fmt.Println("lib2. init() ...")
}
