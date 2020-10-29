package Utils

import (
	"fmt"
	"testing"
)

func tostring(data ASTNode)  {
	data.Tostring()
}

func TestNumASTNode_Tostring(t *testing.T) {
	var num = NumASTNode{Data: 1}
	fmt.Println(num.Tostring())
}

func TestOperatorASTNode_Tostring(t *testing.T) {
	var name = OperatorASTNode{
		Operator: "+",
		Lnode: NumASTNode{Data: 9},
		Rnode: NumASTNode{Data: 12},
	}
	fmt.Println(name.Tostring())
}