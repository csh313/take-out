package userService

import (
	"errors"
	"gorm.io/gorm"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/model"
	util "hmshop/utils"
	"time"
)

type UserService struct{}

func (s UserService) Login(code req.UserLoginDTO) (*model.User, error) {
	//获取用户openid
	openid, err := util.GetOpenID(code.Code)
	if err != nil {
		return nil, err
	}
	if openid == "" {
		return nil, errors.New("openid is null")
	}
	var user model.User
	if err = global.DB.Where("openid = ?", openid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.OpenID = openid
			user.CreateTime = time.Now()
			if err := global.DB.Create(&user).Error; err != nil {
				return nil, err
			} else {

				return &user, nil
			}
		} else {
			return nil, err
		}
	}
	return &user, nil

}
