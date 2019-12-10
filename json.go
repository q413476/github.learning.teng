package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string `json:"type"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type VCard struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Addresses []*Address `json:"addresses"`
	Remark    string     `json:"remark"`
}

type str string

func (s *str) Write(b []byte) (n int, err error) {

	*s = str(b)

	return len(b), err
}

/*
bool 对应 JSON 的 booleans
float64 对应 JSON 的 numbers
string 对应 JSON 的 strings
nil 对应 JSON 的 null
*/
func main() {

	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	js, _ := json.Marshal(vc)

	fmt.Printf("JSON format: %s\n", js)

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)

	defer file.Close()

	var aa str
	//需要一个io.Writer接口参数， file实现了io.Writer
	//返回Encoder
	//enc := json.NewEncoder(file) //存储在文件里面
	enc := json.NewEncoder(&aa) //存储在文件里面

	err := enc.Encode(vc)

	if err != nil {

		log.Println("Error in encoding json")
	}

	fmt.Println(aa)

	//定义一个未知的json字符串
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia",1]}`)

	var f interface{}

	err1 := json.Unmarshal(b, &f)

	fmt.Println(err1)

	fmt.Println(f)

	if m, ok1 := f.(map[string]interface{}); ok1 {
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)

			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don’t know how to handle")
			}
		}

	} else {
		fmt.Println("不是map[string]interface{}")
	}

	jsonStr := `[[1,2,3,4],"aaaa"]`

	var f2 interface{}

	json.Unmarshal([]byte(jsonStr), &f2)

	fmt.Println(f2)

	fmt.Printf("%T\n", f2)

	if m2, ok := f2.(map[string]interface{}); ok {

		fmt.Println(m2)
	} else {
		fmt.Println("f2不是map[string]interface{}")

		if m3, ok2 := f2.([]interface{}); ok2 {

			for kk2, vv2 := range m3 {

				fmt.Println(kk2, vv2, fmt.Sprintf("%T", vv2))
			}

		} else {

			fmt.Println("f2不是[]interface{}")
		}
	}

}
