package router

import (
	"Driving-school/api/v1"
	"Driving-school/middleware"
	"github.com/gin-gonic/gin"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("email").Use(middleware.OperationRecord())
	{
		UserRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
