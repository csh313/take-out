package userController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/api/resp"
	"hmshop/internal/service/userService"
	util "hmshop/utils"
)

type UserApi struct {
	service userService.UserService
}

func (u UserApi) Login(c *gin.Context) {
	var wechatCode req.UserLoginDTO
	if err := c.ShouldBindJSON(&wechatCode); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}

	user, err := u.service.Login(wechatCode)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.LoginError, c)
		return
	}
	//生成jwt令牌
	token, err := util.GenerateToken(c, uint64(user.ID), global.AppConfig.Jwt.User.Name, global.AppConfig.Jwt.User.Secret)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.TokenError, c)
		return
	}
	resq := resp.UserLoginVO{ID: user.ID,
		OpenID: user.OpenID,
		Token:  token}
	fmt.Println(resq)
	res.OkWithData(resq, c)
}
