package main

import (
	"fmt"
)

// 我们一般数据类型写在后边
//
//	下面的返回值是int [如果没有返回值的话，我们直接省略那里就行！]
func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 多返回值情况1：【返回值只写类型，即匿名的】
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

// 多返回值情况2：【返回值 有类型+名称的】
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---- foo3 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//【注意下面的现象】其实这样的是具名返回值，在写返回值名字的时候就已经初始化了
	//r1 r2 属于foo3的形参，  初始化默认的值是0
	//r1 r2 作用域空间 是foo3 整个函数体的{}空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	//whoway：注意，这种返回值写了名称的的，如果你不是显示的return 1000,2000啥的这种通用方式返回
	//那你像上面一样，进行赋值，最后要记得写一下return
	return
}

// 【注意】多返回值情况2的简写： 如果返回值都是int，我们省略的写法(r1,r2 int)
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---- foo4 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	//注意，下面的收方式为 变量1,变量2
	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo4("foo4", 444)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
}
