package v1

import (
	"Driving-school/global"
	"Driving-school/model"
	"Driving-school/model/request"
	"Driving-school/model/response"
	"Driving-school/service"
	"Driving-school/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Coach
// @Summary 创建Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coach true "创建Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coach/createCoach [post]
func CreateCoach(c *gin.Context) {
	Register(c)
}

// @Tags Coach
// @Summary 删除Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coach true "删除Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /coach/deleteCoach [delete]
func DeleteCoach(c *gin.Context) {
	var coach model.Coach
	_ = c.ShouldBindJSON(&coach)
	if err := service.DeleteCoach(coach); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Coach
// @Summary 批量删除Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /coach/deleteCoachByIds [delete]
func DeleteCoachByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteCoachByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Coach
// @Summary 更新Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coach true "更新Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /coach/updateCoach [put]
func UpdateCoach(c *gin.Context) {
	var (
		coach model.Coach
	)
	_ = c.ShouldBindJSON(&coach)

	if err := service.UpdateCoach(coach); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Coach
// @Summary 用id查询Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coach true "用id查询Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /coach/findCoach [get]
func FindCoach(c *gin.Context) {
	var coach model.Coach
	_ = c.ShouldBindQuery(&coach)
	if err, recoach := service.GetCoach(coach.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recoach": recoach}, c)
	}
}

// @Tags Coach
// @Summary 分页获取Coach列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CoachSearch true "分页获取Coach列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coach/getCoachList [get]
func GetCoachList(c *gin.Context) {
	var pageInfo request.CoachSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetCoachInfoList(pageInfo); err != nil {
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
