package main

import "fmt"

//接口提供了一种方式来 说明 对象的行为
// 接口的名字由方法名加 [e]r 后缀组成，例如Printer、Reader、Writer、Logger、Converter 等等。
// 还有一些不常用的方式（当后缀 er 不合适时），
// 比如Recoverable，此时接口名以 able 结尾，或者以 I 开头

//Go 语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法。

//指向接口值的指针是非法的，它们不仅一点用也没有，还会导致代码错误。

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {

	sq1 := new(Square)
	sq1.side = 3.5

	fmt.Println(sq1.Area())

	var sq2 Shaper

	sq2 = sq1

	fmt.Println(sq2.Area())

	r := Rectangle{5, 3}
	q := &Square{5}

	shapes := []Shaper{r, q}

	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	var sq4 Shaper

	sq3 := new(Square)

	sq3.side = 5

	sq4 = sq1

	if t, ok := sq4.(*Square); ok {
		fmt.Printf(" %T\n", t)
	}

	if u, ok := sq4.(*Rectangle); ok {
		fmt.Printf("%T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	switch t := sq4.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Rectangle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
