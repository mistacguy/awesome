package models

// mysql表映射struct
type User struct {
	Id           uint64 `gorm:"primaryKey"`
	Username     string `gorm:"username"`
	Password     string `gorm:"password"`
	status       uint   `gorm:"status"`
	CaptchaId    string `gorm:"captchaId"`
	CaptchaValue string `gorm:"captchaValue"`
}

// 后台管理前端，用户登录参数模型
type WebUserLoginParam struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaId    string `json:"captchaId"`
	CaptchaValue string `json:"captchaValue"`
}

// 后台管理前端，用户信息传输模型
type WebUserInfo struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}

// 微信小程序，用户登录凭证校验模型
type AppCode2Session struct {
	Code      string
	AppId     string
	AppSecret string
}

// 微信小程序，凭证校验后返回的JSON数据包模型
type AppCode2SessionJson struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    uint   `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
