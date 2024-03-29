package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var m, mTimer, mStatus, mFileNum sync.Mutex

var wg sync.WaitGroup

const CMAX = 1000

//记录协程调用次数和文件查询次数
type Perform struct {
	status  bool
	timer   int
	fileNum int
}

/**
状态管理
 */
func (p *Perform) change(b bool) {

	mStatus.Lock()

	defer mStatus.Unlock()

	p.status = b
}

/**
统计协程次数
 */
func (p *Perform) addTimer() {

	mTimer.Lock()

	defer mTimer.Unlock()

	p.timer = p.timer + 1
}

/**
统计查到文件次数
 */
func (p *Perform) addFileNum() {

	mFileNum.Lock()

	defer mFileNum.Unlock()

	p.fileNum = p.fileNum + 1
}

func (p Perform) getStatus() bool {

	return p.status
}

func (p Perform) getTimer() int {

	return p.timer
}

func (p Perform) getFileNum() int {

	return p.fileNum
}

var perform Perform

var wait sync.WaitGroup

//按列读取文件中的数据
func main() {

	t1 := time.Now() // get current time

	ch := make(chan int, CMAX)

	//启动一个协程+1
	wait.Add(1)

	go fileList("../../", ch, true)

	wait.Wait()

	fmt.Println(perform)

	elapsed := time.Since(t1)

	fmt.Println("结束", elapsed)

}

func fileList(path string, c chan int, isGo bool) {

	//是否是协程启动的
	if isGo {

		writeC(c)

		defer releaseC(c)
		//一个协程-1
		defer wait.Done()
	}

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
			//如果当前打开文件小于1000个就交个协程完成，满了就自己完成

			/*if len(c) < CMAX {



			} else {

				fileList(path+val.Name()+"/", c, false)
			}*/
			//启动一个协程+1
			wait.Add(1)

			go fileList(path+val.Name()+"/", c, true)

		} else {

			perform.addFileNum()

			writeFile(path, val)
		}

	}

}
func releaseC(c chan int) {

	<-c
}

func writeC(c chan int) {

	c <- 1

	perform.addTimer()

}

//写入文件
func writeFile(path string, info os.FileInfo) {

	m.Lock()

	defer m.Unlock()

	open, _ := os.OpenFile("./input2.txt", os.O_APPEND|os.O_WRONLY, 0644)

	defer open.Close()

	open.WriteString(path + info.Name() + "\n")

}
