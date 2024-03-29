package main

import "fmt"

//程序在内存中存储它的值，每个内存块（或字）有一个地址，通常用十六进制数表示，如：0x6b0820 或0xf84001d7f0

//Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

//一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，
//在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关

//在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。
//使用一个指针引用一个值被称为间接引用。
//符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值；
// 这被称为反引用（或者内容或者间接引用）操作符；另一种说法是指针转移。

func main() {

	var i1 = 5

	fmt.Printf("值: %d, 它的内存中的位置: %p\n", i1, &i1)

	//声明一个int的指针类型
	var intP *int

	//当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	fmt.Println(intP)

	intP = &i1
	fmt.Println(intP)
	fmt.Println(*intP)

}
