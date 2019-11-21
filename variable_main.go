/**
	基本类型的运用
 */
package main

import "fmt"

func main() {

	var a string //单个变量定义

	var b, c int //多个变量定义

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

	fmt.Println(a, b, c, d, e, f, *g, h, i, j, k)

}

//声明变量的一般形式是使用 var 关键字：var identifier type。
//但如果你的全局变量希望能够被外部包所使用，则需要将首个单词的首字母也大写
//当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil
// 记住，所有的内存在 Go 中都是经过初始化的。
