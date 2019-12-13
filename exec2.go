package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

//func Command(name string, arg ...string) *Cmd
//方法返回一个*Cmd， 用于执行name指定的程序(携带arg参数)
//func (c *Cmd) Run() error
//执行Cmd中包含的命令，阻塞直到命令执行完成
//func (c *Cmd) Start() error
//执行Cmd中包含的命令，该方法立即返回，并不等待命令执行完成
//func (c *Cmd) Wait() error
//该方法会阻塞直到Cmd中的命令执行完成，但该命令必须是被Start方法开始执行的
//func (c *Cmd) Output() ([]byte, error)
//执行Cmd中包含的命令，并返回标准输出的切片
//func (c *Cmd) CombinedOutput() ([]byte, error)
//执行Cmd中包含的命令，并返回标准输出与标准错误合并后的切片
//func (c *Cmd) StdinPipe() (io.WriteCloser, error)
//返回一个管道，该管道会在Cmd中的命令被启动后连接到其标准输入
//func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
//返回一个管道，该管道会在Cmd中的命令被启动后连接到其标准输出
//func (c *Cmd) StderrPipe() (io.ReadCloser, error)
//返回一个管道，该管道会在Cmd中的命令被启动后连接到其标准错误


func main() {

	cmd := exec.Command("/Applications/MAMP/bin/php/php7.1.20/bin/php", "a.php")

	stdin, _ := cmd.StdinPipe()

	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}

	stdin.Write([]byte("go text for grep\n"))

	stdin.Write([]byte("go test text for grep\n"))

	stdin.Close()

	fmt.Println(stdin)

	out_bytes, _ := ioutil.ReadAll(stdout)

	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return
	}

	fmt.Println("Execute finished:" + string(out_bytes))

}
