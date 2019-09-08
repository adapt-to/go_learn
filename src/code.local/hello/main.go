package main // 包归属 package为Go语言的关键字：用于定义当前代码所属包
// main包为特殊包名，它表示当前是一个可执行程序，而不是一个库
import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!你好世界！")
}