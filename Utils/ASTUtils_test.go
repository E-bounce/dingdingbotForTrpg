package Utils

import (
	"fmt"
	"testing"
)

func TestAST_ParseExpression(t *testing.T) {
	expression := "1d4+6*(1+2)"
	tokens,err := Parse(expression)
	if err!=nil {
		fmt.Println(err)
	}
	ast := createAST(tokens,expression)
	if ast.Err!=nil {
		fmt.Println(ast.Err)
	}
	astTree := ast.ParseExpression()
	if ast.Err !=nil {
		fmt.Println(ast.Err)
	}
	fmt.Printf("AST: %+v\n",astTree)
}