package main

import "fmt"

type MyInterface interface {
	// 空接口
}

func Foo(x interface{}) {
	if nil == x {
		fmt.Println("x is a empty interface")
	} else {
		fmt.Println("x is non-empty interface")
	}
}
func main() {
	//这个是interface本身的type和data都是nil, 所以interface[整体为nil]
	Foo(nil)
	var p *int = nil
	//这个是interface本身的type为上面的类型之类的, 只是data为nil, 所以interface[整体不为nil]
	Foo(p)
}
