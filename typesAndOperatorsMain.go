package main

import "fmt"

//只有两个类型相同的值才可以和二元运算符结合
//
//Go 是强类型语言，因此不会进行隐式转换，任何不同类型之间的转换都必须显式说明
//
//如果值的类型是接口，它们也必须都实现了相同的接口
//
//在格式化输出时，你可以使用 %t 来表示你要输出的值为布尔型。
//
//布尔值（以及任何结果为布尔值的表达式）最常用在条件结构的条件语句中，例如：if、for 和 switch 结构
//
//对于布尔值的好的命名能够很好地提升代码的可读性，例如以 is 或者 Is 开头的isSorted、isFinished、isVisivle
//
//Go 语言中没有 float 类型
func main() {

	var a int
	var b int32
	a = 15
	//b = a + a     编译错误
	b = b + 5 // 因为 5 是常量，所以可以通过编译

	fmt.Println(a, b)

	var n int16 = 34
	var m int32

	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
}
