/**
	基本类型的运用
 */
package main

import (
	"fmt"
	"math"
	"os"
)

var Pi float64

//声明变量的一般形式是使用 var 关键字：var identifier type。
//
//但如果你的全局变量希望能够被外部包所使用，则需要将首个单词的首字母也大写
//
//当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil
//
//记住，所有的内存在 Go 中都是经过初始化的。
//
//init函数不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。
//
//每一个源文件都可以包含且只包含一个 init 函数。初始化总是以单线程执行，并且按照包的依赖关系顺序执行。
//
//一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。
func init() {

	Pi = 4 * math.Atan(1)
}

func main() {

	var a string //单个变量定义

	var b, c int //多个变量定义

	a1, b1, c1 := 5, 7, 8

	a1, b1 = b1, c1

	var (
		d int
		e bool
		f string
		g *int
		h *int
		k = "我是预言家"
	)

	var i string = "我是好人" //显形赋值

	var j = "我是女巫" //隐形负责

	a = "我是卧底"

	b, c = 1, 2 //变量进行赋值

	g = &c

	*g = 77777

	fmt.Println(a, b, c, d, e, f, *g, h, i, j, k, a1, b1, c1)

	//不过自动推断类型并不是任何时候都适用的，当你想要给变量的类型并不是自动推断出的某种类型时，你还是需要显式指定变量的类型
	var aa int64 = 2

	var (
		HOME   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)

	fmt.Println(aa, HOME, USER, GOROOT)

	var goos string = os.Getenv("GOOS")

	fmt.Println(goos)

	path := os.Getenv("PATH")

	fmt.Println(path)

	//path := "测试"  再次定义赋值会报错

	fmt.Println(Pi)

}
