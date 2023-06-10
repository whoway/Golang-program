package main //程序的包名

/*
import "fmt"
import "time"
*/
import (
	"fmt"
	"time"
)

// main函数
func main() { //函数的{  一定是 和函数名在同一行的，否则编译错误
	//golang中的表达式，加";", 和不加 都可以，建议是不加, 因为标准库的代码也没加
	fmt.Println(" hello Golang!")
	time.Sleep(1 * time.Second)
}
