package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lemonsoul/gin-demo/common"
)

func InitRouter() {
	Router := gin.Default()
	Router.LoadHTMLGlob("templates/*")
	base := Router.Group(baseUrl)
	Base(base)
	err := Router.Run(common.PORT)
	if err != nil {
		fmt.Println("启动失败！")
	}
}
