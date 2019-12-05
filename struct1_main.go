package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//首写字母为小写，不会json转换
type T struct {
	F1 int
	F2 int
	F3 int
	f4 int
}

type T1 struct {
	f string `one:"1" two:"2"blank:""`
}

type peerInfo struct {
	HTTPPort int    `json:"http_port"`
	TCPPort  int    `json:"tcp_port"`
	SSHPort  int    `json:"ssh_port"`
	Pwd      string `json:"-"`
	versiong string
}

func main() {

	t := new(T1)
	t2 := reflect.TypeOf(t)
	fmt.Println(t2)
	fmt.Println(t2.Elem().Kind())
	fmt.Println(t2.Elem().FieldByName("f"))
	fmt.Println(t2.Elem().NumField())
	fmt.Println(t2.Elem().Field(0).Tag.Get("one"))
	fmt.Println(t2.Elem().Field(0).Tag.Get("two"))

	fmt.Println("----------", "json")

	t3 := T{1, 2, 3, 4}

	t4 := [] T{t3, t3, t3, t3, t3, t3}

	b, err := json.Marshal(t4)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

	fmt.Println("----------", "struct=>json", "tag")

	pi := peerInfo{80, 3306, 22, "我是密码", "0.0.1"}
	js, err := json.Marshal(pi)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(js)
	fmt.Println(string(js))

	fmt.Println("----------", "json=>struct")

	var v peerInfo
	data := []byte(`{"http_port":80,"tcp_port":3306}`)
	err3 := json.Unmarshal(data, &v)
	if err3 != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", v)

}
