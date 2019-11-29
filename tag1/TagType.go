package tag1

import (
	"fmt"
	"reflect"
)

//一个带标签的结构体
type TagType struct {
	Field1 bool   "An important answer"
	Field2 string "The name of the thing"
	Field3 int    "How much there are"
}

func (t TagType) TefTag() {

	ttType := reflect.TypeOf(t)
	for i := 0; i < 3; i++ {
		ixField := ttType.Field(i)
		fmt.Printf("%v\n", ixField.Tag)
	}
}
