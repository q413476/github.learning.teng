package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//反射是用程序检查其所拥有的结构，尤其是类型的一种能力；这是元编程的一种形式。
//反射可以在运行时检查类型和变量，例如它的大小、方法和动态的调用这些方法
//变量的最基本信息就是类型和值：反射包的Type用来表示一个Go类型，反射包的Value为Go值提供了反射接口。
//两个简单的函数,reflect.TypeOf和reflect.ValueOf，返回被检查对象的类型和值

type Kind uint

//为什么需要 Kind =这个操作？ 因为iota默认是int类型，可以定义 Kind来转义特殊类型需要


const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

// String returns the name of k.
func (k Kind) String() string {
	if int(k) < len(kindNames) {
		return kindNames[k]
	}
	return "kind" + strconv.Itoa(int(k))
}

var kindNames = []string{
	Invalid:       "invalid",
	Bool:          "bool",
	Int:           "int",
	Int8:          "int8",
	Int16:         "int16",
	Int32:         "int32",
	Int64:         "int64",
	Uint:          "uint",
	Uint8:         "uint8",
	Uint16:        "uint16",
	Uint32:        "uint32",
	Uint64:        "uint64",
	Uintptr:       "uintptr",
	Float32:       "float32",
	Float64:       "float64",
	Complex64:     "complex64",
	Complex128:    "complex128",
	Array:         "array",
	Chan:          "chan",
	Func:          "func",
	Interface:     "interface",
	Map:           "map",
	Ptr:           "ptr",
	Slice:         "slice",
	String:        "string",
	Struct:        "struct",
	UnsafePointer: "unsafe.Pointer",
}

type flag uintptr

const (
	flagKindWidth        = 5 // there are 27 kinds
	flagKindMask    flag = 1<<flagKindWidth - 1
	flagStickyRO    flag = 1 << 5
	flagEmbedRO     flag = 1 << 6
	flagIndir       flag = 1 << 7
	flagAddr        flag = 1 << 8
	flagMethod      flag = 1 << 9
	flagMethodShift      = 10
	flagRO          flag = flagStickyRO | flagEmbedRO
)

func (f flag) kind() Kind {
	return Kind(f & flagKindMask)
}

func main() {
	var xx = 5.7

	fmt.Println(reflect.TypeOf(xx))
	fmt.Println(reflect.ValueOf(xx))

	type MyInt int
	var m MyInt = 5
	vv := reflect.ValueOf(m)

	fmt.Println(vv.Kind())
	fmt.Println(vv.Interface())

	fmt.Println("--------------")

	var x = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

	fmt.Println("--------------")

	var f flag = 5531233123123

	fmt.Println(f.kind())

}
