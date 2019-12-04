package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//进行排序 冒泡排序
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

//判断切片里面是否有 前置位置比后置位置索引对应值小的
func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

//定义int类型的切片
type IntArray []int

//获取切片的长度
func (p IntArray) Len() int {
	return len(p)
}

//比较索引对应值的大小
func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

//2个索引的值进行切换位置
func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//定义string类型的切片
type StringArray []string

//获取string类型切片的长度
func (p StringArray) Len() int {
	return len(p)
}

//比价2个索引对应字符串的大小
func (p StringArray) Less(i, j int) bool {
	return p[i] < p[j]
}

//2个索引对应字符串进行调换位置
func (p StringArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//int类型切片排序
func SortInts(a []int) {
	Sort(IntArray(a))
}

//string类型切片排序
func SortStrings(a []string) {
	Sort(StringArray(a))
}

func IntsAreSorted(a []int) bool {
	return IsSorted(IntArray(a))
}
func StringsAreSorted(a []string) bool {
	return IsSorted(StringArray(a))
}
