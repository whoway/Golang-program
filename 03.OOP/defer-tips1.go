package main

//知识点2: defer和return谁先谁后
import (
	"fmt"
)

func haveDefer() int {
	fmt.Println("defer func...")
	return 0
}

func notDefer() int {
	fmt.Println("no defer ...")
	return 0
}

func demo() int {
	defer haveDefer()
	return notDefer()
}

func main() {
	demo()
}
