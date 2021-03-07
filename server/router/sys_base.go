package router

import (
	"Driving-school/api/v1"
	"Driving-school/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base").Use(middleware.NeedInit())
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
		BaseRouter.POST("music", v1.GetMusic)
		BaseRouter.POST("getWeather", v1.GetWeather)
		BaseRouter.POST("getDashboardInfo", v1.GetDashboardInfo) // 获取时间信息
	}
	return BaseRouter
}
