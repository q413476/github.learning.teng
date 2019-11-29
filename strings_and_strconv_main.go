package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {

	str1 := "为天地立心，为万灵立命，为往圣继绝学，为万世开太平"

	//是否以某个字符串开头
	is := strings.HasPrefix(str1, "太平")
	fmt.Println(is)

	//是否以某个字符串结束
	is1 := strings.HasSuffix(str1, "太平")
	fmt.Println(is1)

	//是否包含某个字符串
	is2 := strings.Contains(str1, "万灵")
	fmt.Println(is2)

	//获取字符串的第一个字符的索引
	index1 := strings.Index(str1, "，")
	fmt.Println(index1)
	//获取字符串的最后一次索引位置
	index2 := strings.LastIndex(str1, "，")
	fmt.Println(index2)

	index3 := strings.IndexAny(str1, "，")
	fmt.Println(index3)

	//统计出现次数
	fmt.Println(strings.Count(str1, "，"))

	//字符串替换
	fmt.Println(strings.Replace(str1, "，", "-", strings.Count(str1, "，")))

	//重复字符串
	fmt.Println(strings.Repeat(str1, 3))

	str2 := " Today's featured article 1"

	//将字符串改成小写
	fmt.Println(strings.ToLower(str2))

	//将字符串改成大写
	fmt.Println(strings.ToUpper(str2))

	//修剪字符串来剔除字符串开头和结尾的空白符号
	fmt.Println(strings.TrimSpace(str2))

	//剔除指定字符
	fmt.Println(strings.Trim(str2, "1"))

	//通过空白符进行分割
	sl := strings.Fields(str2)

	for _, val := range sl {
		fmt.Printf("%s", val)
	}

	fmt.Println()

	//通过制定字符进行分割
	str3 := " Today's|featured|article"
	sl1 := strings.Split(str3, "|")

	fmt.Println(sl1)

	for _, val1 := range sl1 {
		fmt.Printf("%s", val1)
	}

	fmt.Println()

	//字符串拼接

	sc1 := [] string{"e", "c", "h", "o"}
	fmt.Println(sc1)
	fmt.Println(strings.Join(sc1, "^"))

	//读取内容
	rea := strings.NewReader(str2)

	b := make([]byte, 8)

	for {
		n, err := rea.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	fmt.Println(strconv.IntSize)

	str5 := strconv.Itoa(999)

	fmt.Println(str5)

	fmt.Printf("str5 type:%T\n", str5)

	// em必须为initerface类型才可以进行类型断言

	//if value, ok := str5.(string); ok == true {
	//	fmt.Printf("x[%d] 类型为int,内容为%d\n", index, value)
	//}   cannot type switch on non-interface value str5 (type string)

	if val1, ok := interface{}(str5).(string); ok {
		fmt.Printf("%s:是string类型\n", val1)
	}

	switch val2 := interface{}(str5).(type) {
		case int:
			fmt.Println("int", val2)
		case string:
			fmt.Println("string", val2)
		default:
			fmt.Println("unknown", val2)
	}

	type Element interface{}

	if val1, ok := Element(str5).(string); ok {
		fmt.Printf("%s:是string类型\n", val1)
	}

	i1 := 9.213213123123123213123123

	fmt.Println(i1)

	fmt.Printf("i1 type:%T\n", i1)

	i2 := strconv.FormatFloat(i1, 'b', -1, 64)

	fmt.Println(i2)

	fmt.Printf("i2 type:%T\n", i2)

}
