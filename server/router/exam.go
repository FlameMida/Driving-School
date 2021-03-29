package router

import (
	"Driving-school/api/v1"
	"Driving-school/middleware"
	"github.com/gin-gonic/gin"
)

func InitExamRouter(Router *gin.RouterGroup) {
	ExamRouter := Router.Group("Exam").Use(middleware.OperationRecord())
	{
		ExamRouter.POST("createExam", v1.CreateExam)             // 新建Exam
		ExamRouter.DELETE("deleteExam", v1.DeleteExam)           // 删除Exam
		ExamRouter.DELETE("deleteExamByIds", v1.DeleteExamByIds) // 批量删除Exam
		ExamRouter.PUT("updateExam", v1.UpdateExam)              // 更新Exam
		ExamRouter.GET("findExam", v1.FindExam)                  // 根据ID获取Exam
		ExamRouter.GET("getExamList", v1.GetExamList)            // 获取Exam列表
	}
}
