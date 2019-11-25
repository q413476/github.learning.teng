package main

import (
	"fmt"
)

type Users struct {
	Name    string
	Age     int
	Sex     int
	Address string
}

func main() {

	a, b, _, apple := MinMax(112, 23)

	fmt.Println(a, b, apple)

	fmt.Println(Min(110, 22, 53, 41, 323, 64, 77, 81, 29, 39, ))
	//传切片
	fmt.Println(Min([] int{5, 43, 123, 42, 534, 5, 654, 6, 21}...))

	F1([] string{"a", "b"}...)

	F4([] Users{{"战三", 20, 1, "武汉市"}, {"战三", 20, 1, "武汉市"}}...)

	a1 := new([2] int)

	a3 := [2] int{1, 23}

	a4 := &a3

	//(*a1)[0] = 12312

	a2 := make([] int, 2)

	a2[0] = 12312

	fmt.Println(a1, a2, a3, a4)

	user := new(Users)

	typeCheck(1, "上帝之杖", false, 123123, 9.213213, "世界之窗", user)

	fmt.Println(user)
}

//多类型参数进行断言
func typeCheck(a int, b string, values ...interface{}) {

	defer fmt.Println("我已经结束了", a, b)
	for _, value := range values {
		switch v := value.(type) {
		case int:
			fmt.Println("我是Int类型", v)
		case float64:

			fmt.Println("我是float64类型", v)
		case string:
			fmt.Println("我是string类型", v)
		case bool:
			fmt.Println("我是bool类型", v)
		case *Users:
			v.Name = "dasdasd"
			fmt.Println("我是Users类型", v)
		default:
			fmt.Printf("%T\n", v)
			fmt.Println("不知道什么类型", v)
		}
	}

}


func F4(users ...Users) {

	for _, user := range users {

		fmt.Println(user.Name, user.Address, user.Age, user.Sex)

	}
}

func F1(s ...string) {

	F2(s...)

	F3(s)
}
func F2(s ...string) {

	fmt.Println("我是F2中：", s)
}

func F3(s []  string) {

	fmt.Println("我是F3中：", s)
}

//传递变长参数
func Min(a ...int) int {

	if len(a) == 0 {
		return 0
	}

	min := a[0]

	for _, i := range a {

		if i < min {
			min = i
		}

	}

	return min

}

func MinMax(a int, b int) (min int, max int, zone int, apple string) {
	if a < b {

		min = a
		max = b
	} else {

		min = b
		max = a
	}
	zone = 0
	apple = "我是手机"
	return
}
