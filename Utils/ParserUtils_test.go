package Utils

import (
	"fmt"
	"testing"
)

func TestParser_ParseExpression(t *testing.T) {
	str := "1d20+4"
	a,err := Parse(str)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Printf("data:%v",a)
}