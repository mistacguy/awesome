package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"zshf.private/common"
	"zshf.private/global"
	"zshf.private/models"
	"zshf.private/response"
	"zshf.private/service"
)

var userService service.WebUserService

// WebGetCaptcha 后台管理前端，获取验证码
func WebGetCaptcha(c *gin.Context) {
	id, b64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{"captchaId": id, "captchaImg": b64s}
	response.Success("操作成功", data, c)
}

// WebUserLogin 后台管理前端，用户登录
func WebUserLogin(c *gin.Context) {
	var param models.WebUserLoginParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}

	// 检查验证码
	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	}
	// 生成token
	uid := userService.Login(param)
	if uid > 0 {
		token, _ := common.GenerateToke(param.Username)
		userInfo := models.WebUserInfo{
			Uid:   uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}

// AppUserLogin 微信小程序，用户登录
func AppUserLogin(c *gin.Context) {
	var acsJson models.AppCode2SessionJson
	acs := models.AppCode2Session{
		Code:      c.PostForm("code"),
		AppId:     global.Config.Code2Session.AppId,
		AppSecret: global.Config.Code2Session.AppSecret,
	}
	api := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	path := fmt.Sprintf(api, acs.AppId, acs.AppSecret, acs.Code)
	res, err := http.DefaultClient.Get(path)
	if err != nil {
		fmt.Println("微信登录凭证校验接口请求错误")
		return
	}
	if err := json.NewDecoder(res.Body).Decode(&acsJson); err != nil {
		fmt.Println("decoder error...")
		return
	}
	response.Success("登录成功", acsJson.OpenId, c)
}
