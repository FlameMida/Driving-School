package service

import (
	"Driving-school/global"
	"Driving-school/model"
	"Driving-school/model/request"
)

//@function: CreateExam
//@description: 创建Exam记录
//@param: Exam model.SysExam
//@return: err error

func CreateExam(Exam model.SysExam) (err error) {
	err = global.GVA_DB.Create(&Exam).Error
	return err
}

//@function: DeleteExam
//@description: 删除Exam记录
//@param: Exam model.SysExam
//@return: err error

func DeleteExam(Exam model.SysExam) (err error) {
	err = global.GVA_DB.Delete(&Exam).Error
	return err
}

//@function: DeleteExamByIds
//@description: 批量删除Exam记录
//@param: ids request.IdsReq
//@return: err error

func DeleteExamByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SysExam{}, "id in ?", ids.Ids).Error
	return err
}

//@function: UpdateExam
//@description: 更新Exam记录
//@param: Exam *model.SysExam
//@return: err error

func UpdateExam(Exam model.SysExam) (err error) {
	err = global.GVA_DB.Omit("password").Updates(&Exam).Error
	return err
}

//@function: GetExam
//@description: 根据id获取Exam记录
//@param: id uint
//@return: err error, Exam model.SysExam

func GetExam(id uint) (err error, Exam model.SysExam) {
	err = global.GVA_DB.Where("id = ?", id).First(&Exam).Error
	return
}

//@function: GetExamInfoList
//@description: 分页获取Exam记录
//@param: info request.ExamSearch
//@return: err error, list interface{}, total int64

func GetExamInfoList(info request.ExamSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysExam{})
	var Exams []model.SysExam
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId > 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Exams).Error

	return err, Exams, total
}

//@function: GetExamDetailList
//@description: 获取某个学员的Exam记录
//@param: info request.ExamSearch
//@return: err error, list interface{}, total int64

func GetExamDetailList(info request.ExamSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysExam{}).Where("user_id = ?", info.UserId)
	var Exams []model.SysExam
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Exams).Error

	return err, Exams, total
}
