package tool

import (
	"fmt"
	"testing"
)

func Test_Struct(t *testing.T) {
	type TEST struct {
		TestStr   string
		TestInt   int
		TestFloat float64
		TestBool  bool
	}
	teststruct := TEST{"test", 2, 0.6, true}
	m := StructToMap(teststruct)
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
}
