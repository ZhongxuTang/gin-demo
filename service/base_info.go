package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonsoul/gin-demo/common"
	"github.com/lemonsoul/gin-demo/config"
	"github.com/lemonsoul/gin-demo/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

type BaseApi struct{}

func (baseApi *BaseApi) GetUserInfo(context *gin.Context) {
	testParam := context.Query("test")
	logrus.Infoln("query param ", testParam)
	var user = make([]*model.User, 10)
	userResult := config.DB.Find(&user)
	userResult.Row()
	logrus.Infoln("test result: ", user)
	common.SuccessWithData(gin.H{
		"list": user,
	}, context)
}

func (baseApi *BaseApi) SetUserInfo(context *gin.Context) {
	common.Success(context)
}

func (baseApi *BaseApi) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "lemonsoul",
	})
}
