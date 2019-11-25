package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {

	f()

	b()

	f2("this i")

	a := callback(23, func(i int, i2 int) int {

		return i * i2
	})

	fmt.Println(a)
	fmt.Println("-----------")
	aaa := strings.IndexFunc("this is good", func(r rune) bool {
		return rune('d') == r
	})
	fmt.Println(aaa)

	fun1 := func(i int, j int) int {

		return i + j
	}

	func(i int, j int) {
		fmt.Println(i, j)
	}(23, 321312)

	fmt.Println(fun1(1, 4))
	fmt.Println("-----")
	fmt.Println(f4())
}

func f4() (ret int) {

	defer func() {
		ret++
	}()
	return 1

}

func callback(i int, f func(int, string int) int) int {

	return i + f(i, 3)
}

//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出
func f() {

	for i := 0; i < 5; i++ {

		defer fmt.Println(i)
	}

}

func trace(s string) string {
	fmt.Println("进入:", s)
	return s
}

func un(s string) {
	fmt.Println("离开:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("在 a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("在 b")
	a()
}

func f2(s string) (n int, err error) {

	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()

	return 7, io.EOF

}
