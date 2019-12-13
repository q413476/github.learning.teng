package main

import (
	"fmt"
	"os"
	"sync"
)

const CMAX3 = 1000

var wait3 sync.WaitGroup

//对fileinput2.go进行优化
func main() {

	ch := make(chan string, CMAX3)

	wait3.Add(1)
	go getFileList("../../", ch)

	for va := range ch {
		fmt.Println(va)
	}
	wait3.Wait()

}

func getFileList(path string, c chan string) {

	//打开一个目录
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {

		panic(err)
	}

	//读取目录下面所有文件以及目录
	a, err := file.Readdir(1024)

	for _, val := range a {

		if val.IsDir() {

			wait3.Add(1)
			
			defer wait3.Done()
			
			go getFileList(path+val.Name()+"/", c)

		} else {

			c <- path + val.Name()
		}

	}
}
