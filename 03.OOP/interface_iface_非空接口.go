package main

import "fmt"

type MyInterface interface {
	Show()
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
		fmt.Println("Myterface is not nil, but iface's data is nil")
	}
}
