package main

import (
	"fmt"
	"github.learning.teng/file"
	"github.learning.teng/inner1"
	"github.learning.teng/intFunc1"
	"github.learning.teng/person"
	"github.learning.teng/point"
	"github.learning.teng/tag1"
	"unsafe"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func setStr(s *struct1, str string) {
	s.str = str
}
func main() {

	//2种不同方法实例化结构体
	ms := new(struct1)
	fmt.Println(ms)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)

	ms1 := &struct1{10, 15.5, "Chris"}
	fmt.Println(ms1)
	setStr(ms1, "bbbbb")
	fmt.Println(ms1)

	//结构体工厂
	//var file file.file = file.NewFile(1, "./text.txt")

	file := file.NewFile(1, "./text.txt")

	fmt.Println(file)
	fmt.Printf("%T\n", file)

	fmt.Println("一个实例占用了多少内存", unsafe.Sizeof(file))

	//使用带标签的结构体

	tt := &tag1.TagType{true, "张三", 1}

	tt.TefTag()
	fmt.Println("匿名（或内嵌）字段")

	i1 := new(inner1.OuterS)
	i1.Output()
	i1.B = 13
	i1.C = "张三"
	i1.Int1 = 12
	i1.Int2 = 13
	i1.InnerS.Int1 = 14
	i1.InnerA.Int3 = 16
	i1.InnerA.Name = "李四"
	i1.Int = 115
	i1.Output()
	i1.SetInt3(1111)
	fmt.Println(i1.Int3)
	i1.SetInt3(4444)
	fmt.Println(i1.InnerA.Int3)

	//int类型添加方法

	n := new(intFunc1.Number)

	n.SetVal(2)

	fmt.Println(n.GetVal())

	fmt.Println("方法和未导出字段")

	person := new(person.Person)

	fmt.Println(person.FirstName())
	person.SetFirstName("zhang")
	fmt.Println(person.FirstName())

	fmt.Println("内嵌类型的方法和继承")

	n11 := &point.NamedPoint{point.Point{2, 3}, "时间"}
	fmt.Println(n11.Abs())       // 360.5551275463989
	fmt.Println(n11.Point.Abs()) // 3.605551275463989

}
