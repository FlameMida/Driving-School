package service

import (
	"Driving-school/config"
	"Driving-school/global"
	"Driving-school/model"
	"Driving-school/utils"
	"go.uber.org/zap"
)

//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server

func GetSystemConfig() (conf config.Server, err error) {
	return global.GVA_CONFIG, nil
}

// @description   set system config,

//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func SetSystemConfig(system model.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils.InitDisk() Failed!", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}

//@author: FlameMida
//@function: GetDashboardInfo
//@description: 获取仪表板信息
//@return: info string, msg string

func GetDashboardInfo() (dashboard *utils.DashBoard, err error) {
	var d utils.DashBoard
	d.TimeInfo, d.TimeMsg = utils.InitTimeInfo()

	if d.ActiveUsers, err = utils.InitActiveUserInfo(); err != nil {
		global.GVA_LOG.Error("d.ActiveUsers redis get Failed!", zap.String("err", err.Error()))
		return &d, err
	}

	d.TotalUsers = int(global.GVA_DB.Find(&[]model.SysUser{}).RowsAffected)
	d.TotalStudents = int(global.GVA_DB.Where("authority_id= ?", "200").Find(&[]model.SysUser{}).RowsAffected)
	d.TotalCoaches = int(global.GVA_DB.Where("authority_id= ?", "100").Find(&[]model.SysUser{}).RowsAffected)
	return &d, nil
}
