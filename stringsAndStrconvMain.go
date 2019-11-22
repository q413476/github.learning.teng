package main

import (
	"fmt"
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

	fmt.Println(strings.TrimSpace(str2))
	fmt.Println(strings.Trim(str2, "1"))

}
