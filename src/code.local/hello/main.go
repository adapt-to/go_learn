package main // 包归属 package为Go语言的关键字：用于定义当前代码所属包
// main包为特殊包名，它表示当前是一个可执行程序，而不是一个库
import ( // import表示代码包的导入（！GO语言规则只有用到的包才能导入，导入但不使用则会报错）
	"fmt"
)

func main() {
	fmt.Println("Hello World!你好世界！")
}