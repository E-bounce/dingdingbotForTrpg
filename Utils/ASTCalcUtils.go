package Utils

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func CreateDiceValue(degree int,max int) (int,string){
	var result = 0
	var buffer bytes.Buffer
	randCreater := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<degree;i++{
		temp := randCreater.Intn(max)+1
		tempString := fmt.Sprintf("第%d次D%d投掷结果为: %d\n",i+1,max,temp)
		result += temp
		buffer.WriteString(tempString)
	}
	return result,buffer.String()
}

func CalcASTResult(ASTtree ASTNode) (int) {
	var ldata,rdata int
	switch ASTtree.(type) {
		case OperatorASTNode:
			ast := ASTtree.(OperatorASTNode)
			ldata = CalcASTResult(ast.Lnode)
			rdata = CalcASTResult(ast.Rnode)
			switch ast.Operator {
				case "+":
					return ldata+rdata
				case "-":
					return ldata-rdata
				case "*":
					return ldata*rdata
				case "/":
					if rdata == 0 {
						panic(errors.New(
							fmt.Sprintf("you can divided zero: [%d/%d]", ldata, rdata)))
					}
					return ldata/rdata
				case "d":
					tempdata,_ := CreateDiceValue(ldata,rdata)
					return tempdata
			default:
			}
		case NumASTNode:
			return ASTtree.(NumASTNode).Data
	}
	return 0
}