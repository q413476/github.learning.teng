package main

import (
	"errors"
	"fmt"
)

/*Go 有一个预先定义的 error 接口类型
type error interface {
	Error() string
}*/

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

type Error1 struct {
	name string
	Err  ErrorString
}

func (e *Error1) Error() string {
	return e.name + ": " + e.Err.Error()
}

type Error2 struct {
	name string
	Err  ErrorString
}

func (e *Error2) Error() string {
	return e.name + ": " + e.Err.Error()
}

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}

func main() {

	err := errors.New("我是错误1号")

	fmt.Println(err.Error())

	var e1 interface{} = new(Error1)

	qq, err22 := e1.(Error1)

	if err22 {
		fmt.Println(1, qq)
	} else {
		fmt.Println(2, qq)
	}

	fmt.Printf("%T\n", e1)

	switch err := e1.(type) {
	case *Error1:
		fmt.Println("Error1", err)
	case *Error2:
		fmt.Println("Error2", err)
	default:
		fmt.Println("未知")
	}

	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")

}
