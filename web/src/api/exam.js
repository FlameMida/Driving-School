import service from '@/utils/request'


// @Tags Exam
// @Summary 创建Exam
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam true "创建Exam"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Exam/createExam [post]
export const createExam = (data) => {
    return service({
        url: "/Exam/createExam",
        method: 'post',
        data
    })
}


// @Tags Exam
// @Summary 删除Exam
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam true "删除Exam"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Exam/deleteExam [delete]
export const deleteExam = (data) => {
    return service({
        url: "/Exam/deleteExam",
        method: 'delete',
        data
    })
}

// @Tags Exam
// @Summary 删除Exam
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Exam"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Exam/deleteExam [delete]
export const deleteExamByIds = (data) => {
    return service({
        url: "/Exam/deleteExamByIds",
        method: 'delete',
        data
    })
}

// @Tags Exam
// @Summary 更新Exam
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam true "更新Exam"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Exam/updateExam [put]
export const updateExam = (data) => {
    return service({
        url: "/Exam/updateExam",
        method: 'put',
        data
    })
}


// @Tags Exam
// @Summary 用id查询Exam
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam true "用id查询Exam"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Exam/findExam [get]
export const findExam = (params) => {
    return service({
        url: "/Exam/findExam",
        method: 'get',
        params
    })
}


// @Tags Exam
// @Summary 分页获取Exam列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Exam列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Exam/getExamList [get]
export const getExamList = (params) => {
    return service({
        url: "/Exam/getExamList",
        method: 'get',
        params
    })
}