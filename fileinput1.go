package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println("将整个文件的内容读到一个字符串里：")

	buf, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile("./input1.txt", buf, 0x644)
	if err != nil {
		panic(err.Error())
	}

}
