package main

import "fmt"

type MyInterface interface {
	// 空接口
}

type Student struct {
}

func (this *Student) Show() {

}

func live() MyInterface {
	var stu *Student
	return stu
}

func main() {
	if nil == live() {
		fmt.Println("Myterface is nil")
	} else {
		fmt.Println("Myterface is not nil, but eface's data is nil")
	}
}
