import service from '@/utils/request'

// @Tags SysApi
// @Summary 创建教练
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.coach true "删除教练"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coach/coach [post]
export const createCoach = (data) => {
    return service({
        url: "/coach/coach",
        method: 'post',
        data
    })
}



// @Tags SysApi
// @Summary 更新教练信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.coach true "更新教练信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coach/coach [put]
export const updateCoach = (data) => {
    return service({
        url: "/coach/coach",
        method: 'put',
        data
    })
}


// @Tags SysApi
// @Summary 删除教练
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.coach true "创建教练"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coach/coach [delete]
export const deleteCoach = (data) => {
    return service({
        url: "/coach/coach",
        method: 'delete',
        data
    })
}


// @Tags SysApi
// @Summary 获取单一教练信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.coach true "获取单一教练信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coach/coach [get]
export const getCoach = (params) => {
    return service({
        url: "/coach/coach",
        method: 'get',
        params
    })
}


// @Tags SysApi
// @Summary 获取权限教练列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限教练列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coach/customerList [get]
export const getCoachList = (params) => {
    return service({
        url: "/coach/customerList",
        method: 'get',
        params
    })
}