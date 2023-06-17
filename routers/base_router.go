package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonsoul/gin-demo/service"
)

type BaseRouter struct {
	path string
}

const baseUrl = "/base"

func Base(routerGroup *gin.RouterGroup) {
	{
		baseApi := service.BaseApi{}
		routerGroup.GET(
			"/getBaseInfo",
			baseApi.GetUserInfo,
		)
		routerGroup.POST(
			"/setBaseInfo",
			baseApi.SetUserInfo,
		)
		routerGroup.GET(
			"/index",
			baseApi.Index,
		)

	}
}
