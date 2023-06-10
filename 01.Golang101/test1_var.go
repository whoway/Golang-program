package main

/*
	4种变量的声明方式
*/

import (
	"fmt"
)

// 声明全局变量 类似【局部变量】的方法1、方法2、方法2是可以的
var gA int = 100
var gB = 200

//用方法4来声明全局变量是不允许的！
//【易错点】 := 只能够用在 函数体内来声明
// := 读作【冒等】
//gC := 200

func main() {

	//方法1：声明一个变量 默认的值是0
	//这是个注意点：和cpp不一样，和Java类似
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法2：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法3：在初始化的时候，可以[省去数据类型]，通过值[自动匹配]当前的变量的[数据类型]
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法4：(常用的方法) 省去var关键字，[直接自动匹配]
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	/*
		测试全局的输出
	*/
	// =====
	fmt.Println("gA = ", gA, ", gB = ", gB)

	/*
		声明多个变量的方法！
	*/
	//方法1
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)

	//方法1的变体
	var kk, ll = 100, "Aceld"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	//方法2：多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
