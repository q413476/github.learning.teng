package main

import "fmt"

//一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
//另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，
//则需要用到 rune 类型。rune 类型等价于 int32 类型。
func main() {

	var ch byte = 'A'
	var ch1 byte = '\x41'

	fmt.Println(ch)
	fmt.Println(ch1)

	var ch2 int = '\u0041'
	var ch3 int = '\u03B2'
	var ch4 int = '\U00101234'

	fmt.Printf("%d - %d - %d\n", ch2, ch3, ch4) // integer
	fmt.Printf("%c - %c - %c\n", ch2, ch3, ch4) // character
	fmt.Printf("%X - %X - %X\n", ch2, ch3, ch4) // UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch2, ch3, ch4)   // UTF-8 code point

}
