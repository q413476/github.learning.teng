package main

import (
	"fmt"
	"strconv"
)

func main() {

	var orig = "ABC"

	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	var num1 = 99

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

	x := [2] int{1, 2}
	y := [3] int{6, 2, 3}

	t := 0
	switch a, b := x[1], y[2]; {
	case a < b:
		t = -1
	case a == b:
		t = 0
	case a > b:
		t = 1
	}
	fmt.Println(t)

	k := 6
	//Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，
	//而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本语"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	//ASCII 编码的字符占用 1 个字节，既每个索引都指向不同的字符
	//而非 ASCII 编码的字符（占有 2 到 4 个字节）不能单纯地使用索引来判断是否为同一个字符

	for s1 := range str {

		fmt.Println(string(str[s1]))
	}

	//标签和 goto 语句之间不能出现定义新变量的语句
	i := 1

LABEL1:

	if i < 5 {
		i++
		fmt.Println(i)
		goto LABEL1
	}

LABEL2:

	for ; i < 10; i++ {
		fmt.Println(i)
		continue LABEL2
	}


}
