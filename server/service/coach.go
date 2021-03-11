package service

import (
	"Driving-school/global"
	"Driving-school/model"
	"Driving-school/model/request"
)

//@function: CreateCoach
//@description: 创建Coach记录
//@param: coach model.Coach
//@return: err error

func CreateCoach(coach model.Coach) (err error) {
	err = global.GVA_DB.Create(&coach).Error
	return err
}

//@function: DeleteCoach
//@description: 删除Coach记录
//@param: coach model.Coach
//@return: err error

func DeleteCoach(coach model.Coach) (err error) {
	err = global.GVA_DB.Delete(&coach).Error
	return err
}

//@function: DeleteCoachByIds
//@description: 批量删除Coach记录
//@param: ids request.IdsReq
//@return: err error

func DeleteCoachByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Coach{}, "id in ?", ids.Ids).Error
	return err
}

//@function: UpdateCoach
//@description: 更新Coach记录
//@param: coach *model.Coach
//@return: err error

func UpdateCoach(coach model.Coach) (err error) {
	err = global.GVA_DB.Save(&coach).Error
	return err
}

//@function: GetCoach
//@description: 根据id获取Coach记录
//@param: id uint
//@return: err error, coach model.Coach

func GetCoach(id uint) (err error, coach model.Coach) {
	err = global.GVA_DB.Where("id = ?", id).First(&coach).Error
	return
}

//@function: GetCoachInfoList
//@description: 分页获取Coach记录
//@param: info request.CoachSearch
//@return: err error, list interface{}, total int64

func GetCoachInfoList(info request.CoachSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Coach{})
	var coachs []model.Coach
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`authority_id` = ? ", "100")
	if info.NickName != "" {
		db = db.Where("`nick_name` LIKE ?", "%"+info.NickName+"%")
	}
	if info.Username != "" {
		db = db.Where("`username` LIKE ?", "%"+info.Username+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&coachs).Error
	return err, coachs, total
}
