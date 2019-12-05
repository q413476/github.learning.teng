package main

import (
	"fmt"
	"reflect"
)

type Child1 struct {
	Name     string `json:"name,name1"`
	Grade    int    `1:"成功" 2:"失败"`
	Handsome bool
}

func (c *Child1) SetName(name string) {

	c.Name = name
}

func (c Child1) GetName() string {

	return c.Name
}

//反射解析

//rtype 是所有类型结构体基础元素，
//emptyInterface:包含了 *rtype，unsafe.Pointer2个类型的元素
//iface 有2个类型 *tab ,data  ;*tab _type以及 interfacetype组成。 _type表示具体类型，而 interfacetype则表示具体类型实现的接口类型

func main() {

	children := []Child1{
		{"Ava", 3, true},
		{"Qcrao", 6, false},
	}

	//在里面转换成emptyInterface 并且返回Value的结构体，Value结构体里面有rtype
	cv := reflect.ValueOf(children)
	fmt.Println(cv)

	//在里面转换成emptyInterface 并且返回rtype
	ct := reflect.TypeOf(children)
	fmt.Println(ct)
	//以上可以得出 ValueOf 可以转成 TypeOf， TypeOf 转换不了 ValueOf

	cvt := cv.Type()
	fmt.Println(cvt)
	// cvt 和 vt是等价的 ,下面2个结果一样

	fmt.Println(cvt.Kind())

	fmt.Println(ct.Kind())

	//和上面结果一样，是英文 Value和Type接口都有Kind的方法
	fmt.Println(cv.Kind())

	fmt.Println("--------", "分界线", "Elem")

	fmt.Println(ct.Elem(), "得到值的类型")

	//cv.Elem() 报错，Value类型 只有 Interface和Ptr才能调用这个方法， 现在是 slice类型
	//下面1注释是Elem的方法
	//.Elem()返回也是Type类型 ，返回是子节点的类型
	//下面注释2中的是一个切片类型有2个属性 rtype，elem； elem也是*rtype，但是他表示是他的元素的类型
	//可以得到Elem()是获取子元素的类型：比如children变量 他本身类型是slice，他的子元素类型是Child1

	cte := ct.Elem()

	fmt.Println(cte)
	fmt.Println(cte.Kind())     //他是个结构体  我要看结构体有多少个字段
	fmt.Println(cte.NumField()) //他有3个字段都打印出去看看

	//通过索引查询
	fmt.Println(cte.Field(0)) // {Name  string json:"name,name1" 0 [0] false}
	fmt.Println(cte.Field(1)) // {Grade  int 1:"成功" 2:"失败" 16 [1] false}
	fmt.Println(cte.Field(2)) // {Handsome  bool  24 [2] false}

	//进行索引为0的操作
	ctef0 := cte.Field(0)
	fmt.Println(ctef0.Name, "名称")
	fmt.Println(ctef0.Tag, "标签")
	fmt.Println(ctef0.Type, "类型")
	fmt.Println(ctef0.Offset, "结构中的偏移量，以字节为单位")
	fmt.Println(ctef0.PkgPath, "包的路径")
	fmt.Println(ctef0.Anonymous, "是不是一个匿名")
	fmt.Println(ctef0.Index, "对应索引")

	//直接通过字段名字查找
	fmt.Println(cte.FieldByName("Name"))
	fmt.Println(cte.NumMethod()) //获取方法数量

	if num1 := cte.NumMethod(); num1 > 0 {

		fmt.Println(cte.Method(0)) //{String  func(main.Child1) string <func(main.Child1) string Value> 0}
		ctem := cte.Method(0)
		fmt.Println(ctem.Name)
		fmt.Println(ctem.Index)
		fmt.Println(ctem.PkgPath)
		fmt.Println(ctem.Type)
		fmt.Println(ctem.Func)
	}

	fmt.Println("-----", "开始修改里面内容")

	fmt.Println(cv.Len(), "切片长度")

	cvLen := cv.Len()

	//通过索引获取切片的元素
	fmt.Println("---", children)

	for i := 0; i < cvLen; i++ {

		e := cv.Index(i)

		//获取每个元素结构体里面Name对象 并且进行修改

		//判断能不能修改
		/*if e.FieldByName("Name").CanSet() {

			e.FieldByName("Name").SetString("aName")


		} else {

			fmt.Println(e.FieldByName("Name"), "不能修改")
		}*/

		//通过索引进行查找，并修改
		if e.CanSet() {
			for ii := 0; ii < e.NumField(); ii ++ {
				fmt.Println(e.Field(ii))

				switch e.Field(ii).Kind() {
				case reflect.String:
					//吧字符串进行反转
					e.Field(ii).SetString(reverseString(e.Field(ii).String()))
				case reflect.Int:
					e.Field(ii).SetInt(e.Field(ii).Int() << 4)
				case reflect.Bool:
					if e.Field(ii).Bool() {
						e.Field(ii).SetBool(false)
					} else {

						e.Field(ii).SetBool(true)
					}
				}
			}
		} else {
			fmt.Println("不能进行设置", e)
		}

	}

	fmt.Println("---", children)
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

/*
1.
func (t *rtype) Elem() Type {
	switch t.Kind() {
	case Array:
		tt := (*arrayType)(unsafe.Pointer(t))
		return toType(tt.elem)
	case Chan:
		tt := (*chanType)(unsafe.Pointer(t))
		return toType(tt.elem)
	case Map:
		tt := (*mapType)(unsafe.Pointer(t))
		return toType(tt.elem)
	case Ptr:
		tt := (*ptrType)(unsafe.Pointer(t))
		return toType(tt.elem)
	case Slice:
		tt := (*sliceType)(unsafe.Pointer(t))
		return toType(tt.elem)
	}
	panic("reflect: Elem of invalid type")
}*/

/*
2.
type sliceType struct {
	rtype
	elem *rtype // slice element type
}*/
