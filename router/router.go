package router

import (
	"dingdingbotForTrpg/Utils"
	"github.com/gin-gonic/gin"
)

var GinObj = gin.Default()

/* 1.calc 计算混合表达式
 * 2. dice 普通的投骰子
 */

func GetRouters() *gin.Engine  {
	GinObj.POST("/calc", func(context *gin.Context) {
		err := Utils.ParseBody(context,"calc")
		if err!=nil {
			context.JSON(200,gin.H{
				"Error": err.Error(),
			})
		}
	})
	GinObj.POST("/dice", func(context *gin.Context) {
		err := Utils.ParseBody(context,"dice")
		if err != nil {
			context.JSON(200,gin.H{
				"Error": err.Error(),
			})
		}
	})

	return GinObj
}
