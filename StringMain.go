package main

import (
	"fmt"
	"strings"
)

//字符串是 UTF-8 字符的一个序列

//当字符为 ASCII 码时则占用 1 个字节，其它字符根据需要占用 2-4 个字节

//字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容；更深入地讲，字符串是字节的定长数组。

//Go 支持以下 2 种形式的字面值：
//解释字符串：该类字符串使用双引号括起来，其中的相关的转义字符将被替换
//非解释字符串：该类字符串使用反引号括起来，支持换行

func main() {

	//解释字符串：其中的相关的转义字符将被替换
	str1 := "换\n 行 \r回车 \t \\"

	//解释字符串：该类字符串使用反引号括起来，支持换行
	str2 := `This is a raw string \n`
	fmt.Println(str1, str2, len(str1))

	str1 = "我是卧底，我是预言家，我是女巫"
	str2 = "This is a raw string"

	//需要注意的是，这种转换方案只对纯 ASCII 码的字符串有效。
	//获取字符串中某个字节的地址的行为是非法的
	fmt.Println(str2[0:3], str1[0:2])

	//由于编译器行尾自动补全分号的缘故，加号 + 必须放在第一行。

	str3 := str2 + str1
	fmt.Println(str3)

	str3 += str1
	fmt.Println(str3)

	//strings.Join()  // fmt.Sprintf
	str4 := strings.Join([] string{str1, str2}, "....")
	fmt.Println(str4)

}
