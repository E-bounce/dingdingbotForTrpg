package main

import (
	"dingdingbotForTrpg/router"
)

func main() {
	r := router.GetRouters()
	r.Run("0.0.0.0:45678")
	//如果需要改到其他端口改成0.0.0.0:Port即可
}

