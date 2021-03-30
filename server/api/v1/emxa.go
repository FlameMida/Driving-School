package v1

import (
	"Driving-school/global"
	"Driving-school/model"
	"Driving-school/model/request"
	"Driving-school/model/response"
	"Driving-school/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Exam
// @Summary 创建考试情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam true "创建考试情况"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Exam/createExam [post]
func CreateExam(c *gin.Context) {
	var Exam model.Exam
	_ = c.ShouldBindJSON(&Exam)
	if err := service.CreateExam(Exam); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Exam
// @Summary 删除Exam
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam true "删除Exam"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Exam/deleteExam [delete]
func DeleteExam(c *gin.Context) {
	var Exam model.Exam
	_ = c.ShouldBindJSON(&Exam)
	if err := service.DeleteExam(Exam); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Exam
// @Summary 批量删除Exam
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Exam"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Exam/deleteExamByIds [delete]
func DeleteExamByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteExamByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Exam
// @Summary 更新Exam
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam true "更新Exam"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Exam/updateExam [put]
func UpdateExam(c *gin.Context) {
	var (
		Exam model.Exam
	)
	_ = c.ShouldBindJSON(&Exam)
	if err := service.UpdateExam(Exam); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Exam
// @Summary 用id查询Exam
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam true "用id查询Exam"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Exam/findExam [get]
func FindExam(c *gin.Context) {
	var Exam model.Exam
	_ = c.ShouldBindQuery(&Exam)
	if err, reExam := service.GetExam(Exam.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reExam": reExam}, c)
	}
}

// @Tags Exam
// @Summary 分页获取Exam列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ExamSearch true "分页获取Exam列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Exam/getExamList [get]
func GetExamList(c *gin.Context) {
	var pageInfo request.ExamSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetExamInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Exam
// @Summary 分页某个学员的Exam列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ExamSearch true "分页获取Exam列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Exam/GetExamDetailList [get]
func GetExamDetailList(c *gin.Context) {
	var pageInfo request.ExamSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetExamDetailList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
