package main

import (
	"fmt"
)

// 声明一种行的数据类型 myint， 是int的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

// 【注意】默认是值传递！
func changeBook(book Book) {
	//传递一个book的副本
	book.auth = "666"
}

// 【注意】这样才是引用传递，其实是传递指针
func changeBook2(book *Book) {
	//指针传递
	book.auth = "777"
}

func main() {

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	fmt.Printf("%v\n", book1)
}
