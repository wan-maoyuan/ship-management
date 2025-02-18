package router

import (
	"ship-management/pkg/web/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func registerPMS(engine *gin.RouterGroup) {
	user := engine.Group("/pms")

	user.GET("/first_category", handler.QueryPMSFirstCategory)
	user.GET("/second_category", handler.QueryPMSSecondCategory)
	user.GET("/third_category", handler.QueryPMSThirdCategory)
	user.GET("/fourth_category", handler.QueryPMSFourthCategory)

	logrus.Info("pms 路由注册成功！！！")
}
