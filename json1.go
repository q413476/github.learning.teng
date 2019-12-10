package main

import (
	"encoding/json"
	"fmt"
)

type FamilyMember struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Parents []string `json:"parents"`
}

type FamilyMembe1 struct {
	Name    string
	Age     int
	Parents []string
}

type FamilyMembe2 struct {
	UserName    string   `json:"user_name"`
	UserAge     int      `json:"user_age"`
	UserParents []string `json:"user_parents"`
}

//预先知道结构
func main() {

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)

	var f FamilyMember

	json.Unmarshal(b, &f)

	fmt.Println(f)

	//会不会识别tag

	b1 := []byte(`{"name": "Wednesday", "age": 6, "parents": ["Gomez", "Morticia"]}`)

	var f1 FamilyMember

	json.Unmarshal(b1, &f1)

	fmt.Println(f1) //可以

	//结构体中没有标签 貌似不用区分大小写
	b2 := []byte(`{"name": "Wednesday", "age": 6, "parents": ["Gomez", "Morticia"]}`)

	var f2 FamilyMembe1

	json.Unmarshal(b2, &f2)

	fmt.Println(f2) //一样可以

	b3 := []byte(`{"user_name": "Wednesday", "user_age": 6, "user_parents": ["Gomez", "Morticia"]}`)

	var f3 FamilyMembe2

	json.Unmarshal(b3, &f3)

	fmt.Println(f3) //可以

	//json 包提供 Decoder 和 Encoder 类型来支持常用 JSON 数据流读写。
	//NewDecoder 和 NewEncoder 函数分别封装了 io.Reader 和 io.Writer 接口。
	//数据结构可以是任何类型，只要其实现了某种接口，
	//目标或源数据要能够被编码就必须实现 io.Writer 或 io.Reader 接口。
	//由于 Go 语言中到处都实现了 Reader 和 Writer，因此 Encoder 和 Decoder 可被应用的场景非常广泛，
	//例如读取或写入 HTTP 连接、websockets 或文件

}
