package request

import uuid "github.com/satori/go.uuid"

// User register structure
type Register struct {
	Username    string `json:"username"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'用户'"`
	HeaderImg   string `json:"headerImg" gorm:"default:''"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
	Phone       string `json:"phone" gorm:"default:''"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify info structure
type ChangePWDStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Modify PhoneNumber structure
type ChangePhoneStruct struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	NewPhone string `json:"newphone"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}
