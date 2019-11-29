package main

import (
	"bytes"
	"fmt"
	"strings"
)

//字符串是 UTF-8 字符的一个序列

//当字符为 ASCII 码时则占用 1 个字节，其它字符根据需要占用 2-4 个字节

//字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容；更深入地讲，字符串是字节的定长数组。

//Go 支持以下 2 种形式的字面值：
//解释字符串：该类字符串使用双引号括起来，其中的相关的转义字符将被替换
//非解释字符串：该类字符串使用反引号括起来，支持换行

//双引号是用来表示字符串string，其实质是一个byte类型的数组，单引号表示rune类型

func main() {

	//解释字符串：其中的相关的转义字符将被替换
	str1 := "换\n 行 \r回车 \t \\"

	//解释字符串：该类字符串使用反引号括起来，支持换行
	str2 := `This is a raw string \n`
	fmt.Println(str1, str2, len(str1))

	str1 = "我是卧底，我是预言家，我是女巫"
	str2 = "This is a raw string"

	fmt.Println(fmt.Sprintf("%s%d", "aaa", 1))

	//需要注意的是，这种转换方案只对纯 ASCII 码的字符串有效。
	//获取字符串中某个字节的地址的行为是非法的
	fmt.Println(str2[0:3], str1[0:2])

	//由于编译器行尾自动补全分号的缘故，加号 + 必须放在第一行。

	str3 := str2 + str1

	fmt.Println(str3)

	str3 += str1
	fmt.Println(str3)

	//strings.Join()
	str4 := strings.Join([] string{str1, str2}, "....")
	fmt.Println(str4)

	//bytes.buffer是一个缓冲byte类型的缓冲器存放着都是byte
	// Buffer 就像一个集装箱容器，可以存东西，取东西（存取数据）
	//创建 一个 Buffer （其实底层就是一个 []byte， 字节切片）
	//向其中写入数据 (Write mtheods)
	//从其中读取数据 (Write methods)

	//4种 写
	var b bytes.Buffer
	b.Write([]byte(str1))
	b.Write([]byte("-------"))
	b.Write([]byte(str2))
	fmt.Println(b.String())

	b1 := new(bytes.Buffer)
	b1.Write([]byte(str1))
	b1.Write([]byte("========"))
	b1.Write([]byte(str2))
	fmt.Println(b1.String())

	//使用[]byte作为参数
	b2 := bytes.NewBuffer([]byte("a"))
	b2.Write([]byte("B"))
	fmt.Println(b2.String())

	//用一个string来初始化可读Buffer,并用string的内容填充Buffer.
	b3 := bytes.NewBufferString("我是好人")
	fmt.Println(b3.Len())
	fmt.Println(b3)

	//读
	b4 := bytes.NewBufferString("为天地立心，为万灵立命，为往圣继绝学，为万世开太平")

	byte1 := make([]byte, b4.Len()-3)

	b4.Read(byte1)

	fmt.Println("我还是剩下多少：", b4.String())

	fmt.Println("被读出来：", string(byte1))

	b4.WriteString("哈哈")

	fmt.Println("满血复活：", b4)

}
