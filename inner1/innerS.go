package inner1

import "fmt"

//匿名（或内嵌）字段
type InnerS struct {
	Int1 int
	Int2 int
}

type InnerA struct {
	//Int2 int 他的存在会导致编译器错误
	Name string
	Int3 int
}

func (i *InnerA) SetInt3(num int) {

		i.Int3 = num
}

type Int int

type OuterS struct {
	B    int
	C    string
	Int1 int
	Int
	InnerS
	InnerA
}

//当两个字段拥有相同的名字
//外层名字会覆盖内层名字，这提供了一种重载字段或方法的方式
//如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了
func (o OuterS) Output() {
	fmt.Println(o.B, o.C, o.Int, o.Int1, o.Int2, o.InnerS.Int1, o.InnerS.Int2)
	fmt.Println(o.Name, o.Int3)
}
