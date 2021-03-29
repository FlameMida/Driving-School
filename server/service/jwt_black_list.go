package service

import (
	"Driving-school/global"
	"Driving-school/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(global.GVA_DB.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: username string
//@return: err error, redisJWT string

func GetRedisJWT(username string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(username).Result()
	return err, redisJWT
}

//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: username string
//@return: err error, redisJWT string

func SetRedisJWT(jwt string, username string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(username, jwt, timer).Err()
	return err
}
