package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := 9.545

	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)

	fmt.Println(iv)
	fmt.Println(it)

	if iv.CanSet() {

		fmt.Println("iv", "开始设置")

	} else {

		//执行上面的代码会产生 panic，原因是反射变量 iv不能代表 i本身
		//因为调用 reflect.ValueOf(i)这一行代码的时候，传入的参数在函数内部只是一个拷贝，是值传递
		//所以 vi代表的只是 i的一个拷贝，因此对 vi进行操作是被禁止的

		fmt.Println("iv", "不能设置")
		fmt.Println("type of p:", iv.Type())
		iv1 := reflect.ValueOf(&i)

		if iv1.CanSet() {

			fmt.Println("iv1", "开始设置")

		} else {

			fmt.Println("iv1", "不能设置")
			fmt.Println("type of p:", iv1.Type())

			iv2 := iv1.Elem()

			if iv2.CanSet() {

				//iv2还不是代表 i， iv2.Elem()才真正代表 i，这样就可以真正操作 i了
				fmt.Println("iv2", "开始设置")
				fmt.Println("type of p:", iv2.Type())

				iv2.SetFloat(1231.3213)
				fmt.Println(i)

			} else {

				fmt.Println("iv2", "不能设置")

			}
		}

	}

	//如果想要操作原变量，反射变量 Value必须要 hold 住原变量的地址才行

	fmt.Println("------", "使用切片试试")
	slice := []int{1, 2, 3, 4, 5}

	vSlice := reflect.ValueOf(slice)
	tSlice := reflect.TypeOf(slice)

	fmt.Println("值：", vSlice)
	fmt.Println("类型：", tSlice)
	fmt.Println("容量：", cap(slice))

	vSlice1 := reflect.ValueOf(&slice)

	fmt.Println(vSlice1)

	if vSlice1.CanSet() {

		fmt.Println("vSlice1", "可以修改")

	} else {

		fmt.Println("vSlice1", "不可以修改")

		vSlice2 := vSlice1.Elem()

		if vSlice2.CanSet() {

			fmt.Println("vSlice2", "可以修改", len(slice), cap(slice))

			//长度和容量都得比原来小,而长度又得比容量小于等于
			vSlice2.SetLen(2)
			vSlice2.SetCap(2)
			fmt.Println("vSlice2：", cap(slice))
		} else {

			fmt.Println("vSlice2", "不可以修改")
		}
	}

	fmt.Println("------------------", "查看系统测试")

	xs := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("xs的容量：", cap(xs), "xs的容量：", len(xs))

	vs := reflect.ValueOf(&xs).Elem()

	vs.SetLen(5)

	fmt.Println("vs的容量：", cap(xs), "vs的长度：", len(xs))

	vs.SetCap(6)

	fmt.Println("vs的容量：", cap(xs), "vs的长度：", len(xs))

	vs.SetCap(5)

	fmt.Println("vs的容量：", cap(xs), "vs的长度：", len(xs))

}
