package Utils

import (
	"fmt"
	"testing"
)


func TestGetSignature(t *testing.T) {
	timestamp := 1603951816708
	//fmt.Println(timestamp)
	fmt.Println(GetSignatureCheck(timestamp))
}