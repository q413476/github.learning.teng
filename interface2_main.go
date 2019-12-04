package main

import (
	"fmt"
	"io"
	"os"
)

//******值接收者和指针接收者的区别*****
type Person1 struct {
	age int
}

//值接收者
func (p Person1) howOld() int {
	return p.age
}

//指针接收者
func (p *Person1) growUp() {
	p.age += 1
}

//******值接收者和指针接收者的区别*****

//******值接收者和指针接收者*****

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	p.language = "aaaaaaaaa"
	fmt.Printf("I am debuging %s language\n", p.language)
}

//******值接收者和指针接收者*****

func main() {

	fmt.Println("-------", "值接收者和指针接收者的区别")
	//方法能给用户自定义的类型添加新的行为。它和函数的区别在于方法有一个接收者
	//给一个函数添加一个接收者，那么它就变成了方法。接收者可以是值接收者，也可以是指针接收者
	//值类型既可以调用值接收者的方法，也可以调用指针接收者的方法；指针类型既可以调用指针接收者的方法，也可以调用值接收者的方法。
	//也就是说，不管方法的接收者是什么类型，该类型的值和指针都可以调用，不必严格符合接收者的类型。

	// qcrao 是值类型
	qcrao := Person1{age: 18}

	fmt.Printf("%T\n", qcrao)

	// 值类型 调用接收者也是值类型的方法
	fmt.Println(qcrao.howOld())

	// 值类型 调用接收者是指针类型的方法
	//qcrao.growUp()  //等价 (&qcrao).growUp()

	fmt.Println(qcrao.howOld())

	// ----------------------

	// stefno 是指针类型
	stefno := &Person1{age: 100}

	fmt.Printf("%T\n", stefno)

	// 指针类型 调用接收者是值类型的方法
	fmt.Println(stefno.howOld()) //等价  (*stefno).howOld()

	// 指针类型 调用接收者也是指针类型的方法
	stefno.growUp()
	fmt.Println(stefno.howOld())

	fmt.Println("-------", "值接收者和指针接收者")
	//不管接收者类型是值类型还是指针类型，都可以通过值类型或指针类型调用，这里面实际上通过语法糖起作用的。
	//实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法
	//而实现了接收者是指针类型的方法，不会自动生成对应接收者是值类型的方法

	var c coder = &Gopher{"Go"}
	c2 := c
	fmt.Printf("%T\n", c)
	c.code()
	c.debug()
	c.code()
	c2.code()

	var c1 Gopher = Gopher{"Go"}
	c3 := c1
	fmt.Printf("%T\n", c1)
	c1.code()
	c1.debug()
	c1.code()
	c3.code()

	var r io.Reader

	fmt.Printf("%T\n", r)
	fmt.Println(r)


	tty, err := os.OpenFile("./test.txt", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("错误")
	}
	r = tty
	fmt.Printf("%T\n", tty)
	fmt.Printf("%T\n", r)

	var w io.Writer

	w = r.(io.Writer)

	w.Write([]byte("dsfsdfsdf"))


	var empty interface{}

	empty = w

	fmt.Printf("%T\n", empty)
	fmt.Println(empty)


}
