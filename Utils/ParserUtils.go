package Utils

import (
	"errors"
	"fmt"
)
// 用常量表示Token类型
const (
	Num = 1
	Operator = 0
)

type Token struct {
	Sign string //当前符号
	Type int // 当前符号类型
}

type Parser struct {
	Str string // 需要解析的字符串
	Ch byte //扫描的当前字符
	Index int //扫描的当前位置
	err error
}

//移动至下一个字符
func (parser *Parser)nextChar()error{
	parser.Index++
	if parser.Index < len(parser.Str){
		parser.Ch = parser.Str[parser.Index]
		return nil
	}
	return errors.New("the string has finished searching ")
}
//跳过空白符
func (parser *Parser)skipBlankSign()error{
	var tempError error
	switch parser.Ch {
	case '\r','\t','\n','\v',' ','\f':
		tempError = parser.nextChar()
	}
	return tempError
}

func (parser *Parser)isNum(num byte) bool{
	return '0' <= num && '9'>= num
}

func (parser *Parser)nextToken()*Token{
	if parser.Index >= len(parser.Str) || parser.err != nil{
		return nil
	}
	_ =parser.skipBlankSign()
	//每次获得token前先跳过空白符
	start := parser.Index
	//记录下当前的扫描位置
	var token *Token
	switch parser.Ch {
		case 'd','+','-','*','/','{','}','(',')','[',']':
			token = &Token{
				Sign: string(parser.Ch),
				Type: Operator,
			}
			//parser.Index = start
			//如果为运算符则刷新当前位置往后移一位
			_ = parser.nextChar()
	case '0','1','2','3','4','5','6','7','8','9':
		for parser.isNum(parser.Ch) && parser.nextChar() == nil {
			}
			/*
			 * 该循环为了确保读取数字直到下一个运算符或空白符为止
			 * 比如读取22这个数字，如果不加此循环只会读到2，而不是22。
			 * 这一步循环会将下标拨到数字的下一位即空白符或者运算符
			 */
			token = &Token{
				Sign: parser.Str[start:parser.Index],
				Type: Num,
			}
	default:
		if parser.Ch !=' ' {
			errorString := fmt.Sprintf("unknown Symbol '%v'",string(parser.Ch))
			parser.err = errors.New(errorString)
			}
		}
		return token
}

func (parser *Parser)ParseExpression()[]*Token{
	tokens := make([]*Token,0)
	for{
		token := parser.nextToken()
		if token == nil {
			break
		}
		tokens =append(tokens,token)
	}
	return tokens
}

func Parse(str string)([]*Token,error){
	parser := &Parser{
		Str:   str,
		Ch:    str[0],
		err: nil,
	}
	tokens := parser.ParseExpression()
	if parser.err != nil{
		return nil,parser.err
	}
	return tokens,nil
}