package main

import (
	"fmt"
	"reflect"
)

type C1 struct {
	Name string
	Age  int
}

//int数组进行反射
func main() {

	ii := [2]C1{{"A", 2}, {"B", 4}}

	fmt.Println(ii, "打印")
	iiV := reflect.ValueOf(&ii)
	iiT := reflect.TypeOf(&ii)

	fmt.Println(iiV, "ValueOf")
	fmt.Println(iiT.Kind(), "ValueOf", "Kind")
	fmt.Println(iiT.Elem(), "ValueOf", "Elem")

	fmt.Println(iiT, "TypeOf")
	fmt.Println(iiV.Kind(), "TypeOf", "Kind")
	//fmt.Println(iiV.Elem(), "TypeOf", "Elem")  不在Type Elem2中类型里面

	fmt.Println(iiV.CanSet(), "iiV是否能直接进行修改")

	//想知道里面数量
	//fmt.Println(iiV.NumField(), "iiV的数量") call of reflect.Value.NumField on array Value   NumField是 对结构体里面字段使用的

	fmt.Println(iiV.Kind(), "************")
	fmt.Println(iiV.Elem().Len(), "数组的数量") //************重点
	fmt.Println(iiV.Elem().Kind(), "************")

	fmt.Println(iiT.Elem().Elem().Field(0).Name)

	for i := 0; i < iiV.Elem().Len(); i++ {

		//iiV.Field(ii) 也是对结构体使用的 再看看别的

		fmt.Println(iiV.Elem().Index(i), iiV.Elem().Index(i).Kind(), iiV.Elem().Index(i).Type())

		for j := 0; j < iiV.Elem().Index(i).NumField(); j++ {

			fmt.Println(iiV.Elem().Index(i).Field(j).CanSet(), "是否可以设置")

			switch iiV.Elem().Index(i).Field(j).Kind() {

			case reflect.String:
				iiV.Elem().Index(i).Field(j).SetString("修改")
			case reflect.Int:
				iiV.Elem().Index(i).Field(j).SetInt(99999)

			}

		}
	}

	fmt.Println(ii)
	fmt.Println("0-------------")

	ij := [7]int{1, 2, 3, 4, 5, 6, 7}

	ijV := reflect.ValueOf(&ij)

	fmt.Println(ijV.Kind(), "************")
	ijVe := ijV.Elem()
	fmt.Println(ijVe.Len(), "数组的数量") //************重点
	fmt.Println(ijVe.Kind(), "************")

	for i := 0; i < ijVe.Len(); i++ {

		fmt.Println(ijVe.Index(i))
		fmt.Println(ijVe.Index(i).CanSet())
		ijVe.Index(i).SetInt(88888)
	}

	fmt.Println(ij)

}
