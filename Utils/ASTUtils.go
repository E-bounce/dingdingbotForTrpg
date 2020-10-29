package Utils

import (
	"errors"
	"fmt"
	"strconv"
)

type AST struct {
	Tokens       []*Token //词法分析器的结果
	Expression   string   // 原表达式
	CurrentToken *Token   // 当前Token
	CurrentIndex int      //当前解析的位置
	Err          error    //错误收集
}

var priorityOP = map[string]int{
	"d": 60,
	"*": 40,
	"/": 40,
	"+": 20,
	"-": 20,
}

func createAST(tokens []*Token,str string) *AST{
	ast := &AST{
		Tokens:       tokens,
		Expression:   str,
	}
	if ast.Tokens == nil || len(ast.Tokens)==0 {
		ast.Err = errors.New("empty tokens are given")
	}else {
		ast.CurrentIndex = 0
		ast.CurrentToken = ast.Tokens[0]
	}
	return ast
}

//获取下一个Token
func (ast *AST) getNextToken() *Token {
	ast.CurrentIndex++
	if ast.CurrentIndex < len(ast.Tokens) {
		ast.CurrentToken = ast.Tokens[ast.CurrentIndex]
		return ast.CurrentToken
	}
	return nil
}

//通过Token获得其对应的优先级
func (ast *AST) getPriority() int {
	if priority, ok := priorityOP[ast.CurrentToken.Sign]; ok {
		return priority
	}
	return -1
}

//将数字Token解析为数字节点
func (ast *AST) parseNum() NumASTNode {
	intNum, err := strconv.Atoi(ast.CurrentToken.Sign)
	if err != nil {
		ast.Err = errors.New(
			fmt.Sprintf(
				"%v\nthe symbol must be '(' or '0-9' there,but %s given",
				err.Error(),
				ast.CurrentToken.Sign,
			),
		)
		return NumASTNode{}
	}
	numNode := NumASTNode{Data: intNum}
	ast.getNextToken()
	return numNode
}
//处理圆括号中的表达式
func (ast *AST) parseParenthesis() ASTNode {
	ast.getNextToken()
	expression := ast.ParseExpression()
	if expression == nil {
		return nil
	}
	if ast.CurrentToken.Sign != ")" {
		ast.Err = errors.New(fmt.Sprintf(
			"the symbol must be ')' there,but %s was given",
			ast.CurrentToken.Sign,
		))
		return nil
	}
	ast.getNextToken()
	return expression
}
//解析出一个节点，可能是数字节点也可能是符号节点,解析之后向后移一位
func (ast *AST) parsePriority() ASTNode {
	switch ast.CurrentToken.Type {
	case Num:
		return ast.parseNum()
	case Operator:
		if ast.CurrentToken.Sign == "(" {
			expression:=ast.parseParenthesis()
			return expression
		}else {
			return ast.parseNum()
		}
	default:
		return nil
	}
}
//处理表达式中的符号
func (ast *AST) parseOperator(minPriority int, lnode ASTNode) ASTNode {
	for {
		tokenPrio := ast.getPriority()
		if tokenPrio < minPriority {
			return lnode
		}
		/* 所有数字优先级会被标记成为-1，小于最小优先级0，因此数字节点会直接返回
		 * 所有符号优先级均大于0，因此不走该分支往下走
		 */
		op := ast.CurrentToken.Sign
		if ast.getNextToken() == nil {
			return lnode
		}
		rnode := ast.parsePriority()
		if rnode == nil {
			return nil
		}
		nextPrio := ast.getPriority()
		if tokenPrio < nextPrio {
			rnode = ast.parseOperator(tokenPrio+1,rnode)
			if rnode == nil {
				return nil
			}
		}
		//比较优先级决定
		lnode = OperatorASTNode{
			Operator: op,
			Lnode:    lnode,
			Rnode:    rnode,
		}
	}
}

func (ast *AST) ParseExpression() ASTNode {
	lnode := ast.parsePriority()
	return ast.parseOperator(0,lnode)
}
