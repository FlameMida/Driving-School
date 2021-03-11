package router

import (
	"Driving-school/api/v1"
	"Driving-school/middleware"
	"github.com/gin-gonic/gin"
)

func InitCoachRouter(Router *gin.RouterGroup) {
	CoachRouter := Router.Group("coach").Use(middleware.OperationRecord())
	{
		CoachRouter.POST("createCoach", v1.CreateCoach)             // 新建Coach
		CoachRouter.DELETE("deleteCoach", v1.DeleteCoach)           // 删除Coach
		CoachRouter.DELETE("deleteCoachByIds", v1.DeleteCoachByIds) // 批量删除Coach
		CoachRouter.PUT("updateCoach", v1.UpdateCoach)              // 更新Coach
		CoachRouter.GET("findCoach", v1.FindCoach)                  // 根据ID获取Coach
		CoachRouter.GET("getCoachList", v1.GetCoachList)            // 获取Coach列表
	}
}
