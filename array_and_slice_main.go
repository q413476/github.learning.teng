package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	fmt.Println(b[:], b[:3])

	a := make([]int, 1, 3)

	fmt.Println(a, len(a), cap(a))

	o := make([][]int, 1, 3)
	o1 := [][]int{make([]int, 1)}

	o[0] = make([]int, 2, 4)
	fmt.Println(o)
	fmt.Println(o1)

	//复制
	c := []int{1, 2}
	c1 := []int{3, 4, 5, 6, 7, 8}
	//将c1 索引1到3的值复制到c中，里面有3个元素，但是c长度为2，所以复制是4，5元素
	copy(c, c1[1:3])

	//将c追加的c1中，因为c1容量为6，现在一共有8个元素， c1容量不够，所以在原来的容量上乘以2
	fmt.Println(c, len(c), cap(c), c1, len(c1), cap(c1))

	c1 = append(c1, c...)

	fmt.Println(c1, len(c1), cap(c1))

	//数组转换成切片

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(arr)

	arrSlice := arr[2:5]

	fmt.Println(arrSlice, len(arrSlice), cap(arrSlice), arr, len(arr), cap(arr))

	fmt.Printf("arrSlice类型为：%T;arr的类型为：%T\n", arrSlice, arrSlice)

	fmt.Println("//随着数组里面元素改变，切片里面的元素值也改变")

	arrSlice1 := arrSlice[1:2]

	fmt.Println(arrSlice1, len(arrSlice1), cap(arrSlice1))

	arr[3] = 123

	fmt.Println(arrSlice1, len(arrSlice1), cap(arrSlice1))
	fmt.Println(arrSlice, len(arrSlice), cap(arrSlice))
	fmt.Println(arr, len(arr), cap(arr))

	fmt.Println("随着切片容量改变。不在对原来数组引用")

	arrSlice1 = append(arrSlice1, 3, 4, 5, 6, 7, 8, 98, 9, 23, 12, 3, 12, 3, 13, 12, 3)

	arr[3] = 55555
	fmt.Println(arrSlice1, len(arrSlice1), cap(arrSlice1))
	fmt.Println(arrSlice, len(arrSlice), cap(arrSlice))
	fmt.Println(arr, len(arr), cap(arr))

	fmt.Println("创建一个是原来arrSlice2倍容量的切片,且容量比arr长度还长")

	arrSlice2 := make([]int, (cap(arrSlice))*2, (cap(arrSlice))*2)

	fmt.Println(arrSlice2)

	copy(arrSlice2, arrSlice)

	arr[3] = 77

	fmt.Println(arrSlice2)

	fmt.Println("直接赋值")

	arrSlice3 := make([]int, (cap(arrSlice))*2, (cap(arrSlice))*2)

	fmt.Println(arrSlice3, cap(arrSlice3))

	arrSlice3 = arrSlice[1:]

	arr[3] = 66

	fmt.Println(arrSlice3, cap(arrSlice3))

	fmt.Println("从字符串生成字节切片")

	fmt.Println([]byte("bacdefghijklmnopqrst"))

	cst := make([]byte, 10, 88)
	copy(cst, "abc")
	fmt.Println(cst)

	china := "中国123，个发的鬼地方个12a"

	chinaNumber := []rune(china)
	fmt.Println(chinaNumber)

	chinaNumber1 := []byte(china)
	fmt.Println(chinaNumber1)

	fmt.Println(utf8.RuneCountInString(china))
	fmt.Println(len(china))

	s := "hello"
	fmt.Println(s[1:])

	s1 := []byte(s)
	s1[0] = 'H'
	fmt.Println(string(s1))


}
