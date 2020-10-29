package Utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckValid(c *gin.Context)bool{
	header := c.Request.Header
	_,t_ok := header["Timestamp"]
	_,s_ok := header["Sign"]
	if !t_ok||!s_ok {
		return false
	}
	timestamp,err := strconv.Atoi(header["Timestamp"][0])
	if err!=nil {
		return false
	}
	sign := header["Sign"]
	return CheckAvailable(timestamp,sign[0])
}


func ParseBody(c *gin.Context,routerName string) error {
	if !CheckValid(c) {
		return errors.New("unauthorized error")
	}
	var receiveBody ReceiveData
	err := c.BindJSON(&receiveBody)
	if err != nil {
		return err
	}
	content := receiveBody.Text["content"]
	NickName := receiveBody.SenderNick
	tokens,err := Parse(content)
	if err !=nil {
		return errors.New("expression error!")
	}
	ast := createAST(tokens,content)
	astTree := ast.ParseExpression()
	result := CalcASTResult(astTree)
	switch routerName {
	case "dice":
		Data := SendData{
			Text: map[string]string{
				"content": fmt.Sprintf("@%s 投掷结果为: %d", NickName, result),
			},
			Msgtype: "text",
		}
		c.JSON(200, gin.H{
			"result": Data,
		})
		err := Send(&Data)
		if err!= nil {
			panic(err)
		}
		return nil
	case "calc":
		Data := SendData{
			Msgtype: "text",
			Text: map[string]string{
				"content": fmt.Sprintf("@%s 计算结果为: %s = %d",NickName,content,result),
			},
		}
		c.JSON(200,gin.H{
			"result" : Data,
		})
		err := Send(&Data)
		if err != nil {
			panic(err)
		}
		return nil
	}
	return errors.New("invalid router url!")
}