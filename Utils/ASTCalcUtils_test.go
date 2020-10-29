package Utils

import (
	"fmt"
	"testing"
)

func TestCalcASTResult(t *testing.T) {
	expression := "12d4+6*(1+2)"
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
	fmt.Printf("%s = %d",expression,CalcASTResult(astTree))
}
