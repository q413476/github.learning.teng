package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {

	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {

	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

type specialString string

//指针方法可以通过指针调用
//值方法可以通过值调用
//接收者是值的方法可以通过指针调用，因为指针会首先被解引用
//接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

//类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
//类型 T 的可调用方法集包含接受者为 T 的所有方法
//类型 T 的可调用方法集不包含接受者为 *T 的方法

type Person struct {
	name string
	age  int
}
type Any interface{}

func main() {
	var lst *List = new(List)

	lst.Append(1)

	fmt.Println(lst)

	CountInto(lst, 1, 10)

	fmt.Println(lst)

	var str = "ABC"

	var val Any

	val = 5

	fmt.Printf("%T\n", val)

	val = str

	fmt.Printf("%T\n", val)

	pers1 := new(Person)
	pers1.name = "张三"
	pers1.age = 56

	val = pers1

	fmt.Printf("%T\n", val)

	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type *Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)

	}

	var whatIsThis specialString = "hello"

	testFunc(whatIsThis)

	//定义一个匿名函数的切片,在进行赋值

	var dataSlice [] func() int = make([]func() int, 3)

	dataSlice[0] = func() int {

		return 1
	}
	dataSlice[1] = func() int {

		return 2
	}
	dataSlice[2] = func() int {

		return 3
	}
	fmt.Println(dataSlice)
	fmt.Println(dataSlice[0]())

	//定义一个空接口切片
	var interfaceSlice []interface{} = make([] interface{}, 5)

	fmt.Println(interfaceSlice)

	//再把匿名函数切片值赋值给空接口切片
	for ix, d := range dataSlice {

		interfaceSlice[ix] = d
	}
	//在添加其他几种类型
	interfaceSlice[3] = 4345345
	interfaceSlice[4] = "神气"

	fmt.Println(interfaceSlice)

	//循环断言
	for _, d1 := range interfaceSlice {

		fmt.Printf("%T\n", d1)
		//if _, ok := d1.(func() int); ok {}
		//这里要注意用vv进行强制装换。d1进行装换会报错
		switch vv := d1.(type) {
		case int:
			fmt.Println("我是int类型", int(vv), fmt.Sprintf("%T", vv))
		case func() int:
			fmt.Println("我是func() int类型", func() int(vv)(), fmt.Sprintf("%T", func() int(vv)))
		case string:
			fmt.Println("我是string类型", string(vv), fmt.Sprintf("%T", vv))
		default:
			fmt.Println("不知道什么类型")
		}
	}

}

func testFunc(any interface{}) {
	switch v := any.(type) {
	case bool:
		fmt.Printf("any %v is a bool type\n", v)
	case int:
		fmt.Printf("any %v is an int type\n", v)
	case float32:
		fmt.Printf("any %v is a float32 type\n", v)
	case string:
		fmt.Printf("any %v is a string type\n", v)
	case specialString:
		fmt.Printf("any %v is a special String!\n", v)
	default:
		fmt.Println("unknown type!\n")
	}
}
