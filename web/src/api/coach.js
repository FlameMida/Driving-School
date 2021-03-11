import service from '@/utils/request'

// @Tags Coach
// @Summary 创建Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coach true "创建Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coach/createCoach [post]
export const createCoach = (data) => {
    return service({
        url: "/coach/createCoach",
        method: 'post',
        data
    })
}


// @Tags Coach
// @Summary 删除Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coach true "删除Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /coach/deleteCoach [delete]
export const deleteCoach = (data) => {
    return service({
        url: "/coach/deleteCoach",
        method: 'delete',
        data
    })
}

// @Tags Coach
// @Summary 删除Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /coach/deleteCoach [delete]
export const deleteCoachByIds = (data) => {
    return service({
        url: "/coach/deleteCoachByIds",
        method: 'delete',
        data
    })
}

// @Tags Coach
// @Summary 更新Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coach true "更新Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /coach/updateCoach [put]
export const updateCoach = (data) => {
    return service({
        url: "/coach/updateCoach",
        method: 'put',
        data
    })
}


// @Tags Coach
// @Summary 用id查询Coach
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coach true "用id查询Coach"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /coach/findCoach [get]
export const findCoach = (params) => {
    return service({
        url: "/coach/findCoach",
        method: 'get',
        params
    })
}


// @Tags Coach
// @Summary 分页获取Coach列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Coach列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coach/getCoachList [get]
export const getCoachList = (params) => {
    return service({
        url: "/coach/getCoachList",
        method: 'get',
        params
    })
}