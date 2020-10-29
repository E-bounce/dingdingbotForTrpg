package Utils

import (
	"fmt"
	"strconv"
)

//基础节点同时为了实现多态，为了实现多态下面的接收器就不要用指针了
type ASTNode interface {
	Tostring() string
}

//数字表达式节点
type NumASTNode struct {
	Data int
}

//操作表达式节点
type OperatorASTNode struct {
	Operator string
	Lnode ASTNode
	Rnode ASTNode
}

func (numASTNode NumASTNode) Tostring() string  {
	return fmt.Sprintf("%s",strconv.Itoa(numASTNode.Data))
}

func (operatorASTNode OperatorASTNode) Tostring() string {
	return fmt.Sprintf("(%s %s %s)",
		operatorASTNode.Lnode.Tostring(),
		operatorASTNode.Operator,
		operatorASTNode.Rnode.Tostring(),
	)
}