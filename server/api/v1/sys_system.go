package v1

import (
	"Driving-school/config"
	"Driving-school/global"
	"Driving-school/model"
	"Driving-school/model/response"
	"Driving-school/service"
	"Driving-school/utils"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// @Tags System
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getSystemConfig [post]
func GetSystemConfig(c *gin.Context) {
	if cfg, err := service.GetSystemConfig(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysConfigResponse{Config: cfg}, "获取成功", c)
	}
}

// @Tags System
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /system/setSystemConfig [post]
func SetSystemConfig(c *gin.Context) {
	var sys model.System
	_ = c.ShouldBindJSON(&sys)
	if err := service.SetSystemConfig(sys); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithData("设置成功", c)
	}
}

// 本方法开发中 开发者windows系统 缺少linux系统所需的包 因此搁置
// @Tags System
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"重启系统成功"}"
// @Router /system/reloadSystem [post]
func ReloadSystem(c *gin.Context) {
	if runtime.GOOS == "windows" {
		response.FailWithMessage("系统不支持", c)
		return
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	err := cmd.Run()
	if err != nil {
		global.GVA_LOG.Error("重启系统失败!", zap.Any("err", err))
		response.FailWithMessage("重启系统失败", c)
		return
	}
	response.OkWithMessage("重启系统成功", c)
	return
}

// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getServerInfo [post]
func GetServerInfo(c *gin.Context) {
	if server, err := service.GetServerInfo(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}

}

// @Tags System
// @Summary 获取仪表板信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/getDashboardInfo [post]
func GetDashboardInfo(c *gin.Context) {
	if dashboard, err := service.GetDashboardInfo(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		response.OkWithDetailed(gin.H{"dashboard": dashboard}, "获取成功", c)
	}

}

// @Tags System
// @Summary 获取播放器音乐列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getMusic [post]
func GetMusic(c *gin.Context) {

	if music, err := utils.GetTopList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		response.OkWithDetailed(gin.H{"music": music}, "获取成功", c)
	}

}

// @Tags System
// @Summary 获取天气状况
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/GetWeather [post]
func GetWeather(c *gin.Context) {
	var loc config.Location
	_ = c.ShouldBindJSON(&loc)
	if loc.Log == "" || loc.Lat == "" {
		loc.Log = global.GVA_CONFIG.Weather.Log
		loc.Lat = global.GVA_CONFIG.Weather.Lat
	}
	appKey := global.GVA_CONFIG.Weather.AppKey
	units := global.GVA_CONFIG.Weather.Units
	lang := global.GVA_CONFIG.Weather.Lang
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=%s&lon=%s&appid=%s&lang=%s&units=%s&exclude=%s", loc.Lat, loc.Log, appKey, lang, units, "minutely,hourly,daily,alerts")
	if global.GVA_CONFIG.System.UseMultipoint {
		if global.GVA_REDIS.Exists("weather").Val() > 0 {
			response.OkWithDetailed(gin.H{"weather": global.GVA_REDIS.Get("weather").Val()}, "获取成功", c)
		} else {
			body, err := http.Get(url)
			if err != nil {
				global.GVA_LOG.Error("fail to fetch weather", zap.Any("err", err))
				response.FailWithMessage("获取失败", c)
				return
			} else {
				strings := utils.ProcessBody(body)
				global.GVA_REDIS.Set("weather", strings, time.Minute)
				response.OkWithDetailed(gin.H{"weather": strings}, "获取成功", c)
			}
		}
	} else {
		body, err := http.Get(url)
		if err != nil {
			global.GVA_LOG.Error("fail to fetch weather", zap.Any("err", err))
			response.FailWithMessage("获取失败", c)
			return
		} else {
			strings := utils.ProcessBody(body)
			response.OkWithDetailed(gin.H{"weather": strings}, "获取成功", c)
		}
	}

}
