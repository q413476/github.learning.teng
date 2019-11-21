package main

import "fmt"

//常量使用关键字 const 定义，用于存储不会改变的数据。
//
//存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
//
//常量的定义格式：const identifier [type] = value

//常量的值必须是能够在编译时就能够确定的

//const c2 = getNumber()  报错  const initializer getNumber() is not a constant

//每当 iota 在新的一行被使用时，它的值都会自动加 1；
// 在每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0（ 简单地讲，每遇到一次 const 关键字，iota 就重置为 0

func main() {

	var n int

	a := n + 5 // 无类型的数字型常量 “5” 它的类型在这里变成了 int

	fmt.Println(a)

	const c2, c3, c4 = 0.3232312, 1 / 323213, "diffuses"
	const Ln2 = 0.693147180559945309417232121458 / 176568075500134360255254120680009
	//因为在编译期间自定义函数均属于未知，因此无法用于常量的赋值，但内置函数可以使用，如：len()。
	const Ln2Len = len(c4)
	const (
		Log2E     = 1 / Ln2
		Billion   = 1e9 // float constant
		hardEight = (1 << 100) >> 97
	)
	fmt.Println(c2, c3, Log2E, Billion, hardEight, Ln2Len)

	fmt.Println("-------A")

	const (
		iotaA = iota + 1
		iotaB
		iotaC = 50
		iotaD
		iotaE
		iotaF = iota
		iotaG
		iotaH
		_  //可以使用下划线跳过不想要的值。
		_
		iotaI
		iotaJ
	)

	fmt.Println(iotaA, iotaB, iotaC, iotaD, iotaE, iotaF, iotaG, iotaH, iotaI, iotaJ)

	fmt.Println("-------B定义在一行的情况")

	const (
		Apple, Banana     = iota + 1, iota + 2
		Cherimoya, Durian  // = iota + 1, iota + 2
		Elderberry, Fig    //= iota + 1, iota + 2
	)
	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)

	fmt.Println("-------C位掩码表达式")

	type Allergen int

	const (
		IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
		IgChocolate                         // 1 << 1 which is 00000010
		IgNuts                              // 1 << 2 which is 00000100
		IgStrawberries                      // 1 << 3 which is 00001000
		IgShellfish                         // 1 << 4 which is 00010000
	)

	fmt.Println(IgEggs, IgChocolate, IgNuts, IgStrawberries, IgShellfish)

	fmt.Println("-------D定义数量级")

	type ByteSize float64

	const (
		_           = iota             // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
		MB                             // 1 << (10*2)
		GB                             // 1 << (10*3)
		TB                             // 1 << (10*4)
		PB                             // 1 << (10*5)
		EB                             // 1 << (10*6)
		ZB                             // 1 << (10*7)
		YB                             // 1 << (10*8)
	)
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)

}
