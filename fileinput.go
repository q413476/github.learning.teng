package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	dir1, _ := os.Getwd()
	dir2, _ := os.Executable()

	fmt.Println("当前目录路径1", dir1)
	fmt.Println("当前目录路径2", dir2)

	inputFile, inputError := os.Open("./input.txt")

	if inputError != nil {
		fmt.Println("打开文件保存")
		return
	} else {
		fmt.Println("打开文件成功")
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	buf := make([]byte, 1024*2)

	for {

		//按行读取内容
		//inputString, readerError := inputReader.ReadString('\n')
		//按字节开读取 从文件读到buf中，并返回字节长度
		n, readerError := inputReader.Read(buf)
		if readerError == io.EOF {
			return
		}
		fmt.Println(n)
		fmt.Println(string(buf))
		//	fmt.Printf("文件内容: %s", inputString)
	}

}
