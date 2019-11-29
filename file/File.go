package file

//禁止使用 new 函数，强制用户使用工厂方法，从而使类型变成私有的
//将结构体首写字母改成小写 file

type file struct {
	fd   int    // 文件描述符
	name string // 文件名
}

func NewFile(fd int, name string) *file {

	if fd < 0 {
		return nil
	}
	return &file{fd, name}
}
