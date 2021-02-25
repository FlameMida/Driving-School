import service from '@/utils/request'

// @Tags SysApi
// @Summary 创建学员
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Student true "删除学员"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/student [post]
export const createStudent = (data) => {
    return service({
        url: "/student/student",
        method: 'post',
        data
    })
}



// @Tags SysApi
// @Summary 更新学员信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Student true "更新学员信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/student [put]
export const updateStudent = (data) => {
    return service({
        url: "/student/student",
        method: 'put',
        data
    })
}


// @Tags SysApi
// @Summary 删除学员
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Student true "创建学员"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/student [delete]
export const deleteStudent = (data) => {
    return service({
        url: "/student/student",
        method: 'delete',
        data
    })
}


// @Tags SysApi
// @Summary 获取单一学员信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Student true "获取单一学员信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/student [get]
export const getStudent = (params) => {
    return service({
        url: "/student/student",
        method: 'get',
        params
    })
}


// @Tags SysApi
// @Summary 获取权限学员列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限学员列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/studentList [get]
export const getStudentList = (params) => {
    return service({
        url: "/student/studentList",
        method: 'get',
        params
    })
}