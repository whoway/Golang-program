package main

import (
	"fmt"
)

// 方法2：在Go里面其实可以用const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota， 每行的iota都会累加1, 第一行的iota的默认值是0
	BEIJING  = 10 * iota //iota = 0  所以此处北京=10*0=0
	SHANGHAI             //iota = 1 所以此处上海=10*1=10
	SHENZHEN             //iota = 2 所以此处深圳=10*2=20
)

// 注意下面的坑！
const (
	aaa = 1
	bbb //我没想到bbb和ccc自动也是1
	ccc
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d                      // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f                      // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
	i, k                      // iota = 4, i = iota * 2, k = iota * 3 , i = 8, k = 12
)

func main() {
	// 定义方式1：常量(只读属性)
	const length int = 10

	fmt.Println("length = ", length)

	//length = 100 //常量是不允许修改的。

	fmt.Println("BEIJIGN = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("aaa = ", aaa)
	fmt.Println("bbb = ", bbb)
	fmt.Println("ccc = ", ccc)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota 只能够配合const() 一起使用， iota只有在const进行累加效果。
	// 【注意】iota 只能够配合const() 一起使用的意思是，iota只能出现在【枚举】中
	// 虽然出现在下面这样的语法中也可以，但是iota那你怎么解释意思？
	// 因而我们认为下面的是语法可行！但是意义不明显，因而少用
	//const s int = iota
}
