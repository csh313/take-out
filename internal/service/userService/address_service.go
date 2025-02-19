package userService

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hmshop/common/code"
	"hmshop/common/enum"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/model"
)

type AddressService struct{}

func (s AddressService) AddressList(c *gin.Context) []model.AddressBook {
	value, exists := c.Get(enum.CurrentUserId)
	fmt.Println(value, "-----")
	if !exists {
		global.Log.Error(errors.New("查询用户信息失败"))
		res.FailWithMessage(code.ReqError, c)
	}
	var addressModel []model.AddressBook
	if err := global.DB.Find(&addressModel, "user_id = ?", value).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.SqlError, c)
	}
	fmt.Println(addressModel)
	return addressModel
}

func (s AddressService) GetById(id int, c *gin.Context) (*model.AddressBook, error) {
	var addressModel model.AddressBook
	if err := global.DB.Where("id = ?", id).First(&addressModel).Error; err != nil {
		return nil, err
	}
	return &addressModel, nil

}

func (s AddressService) GetDefaultAddress(value any, c *gin.Context) (*model.AddressBook, error) {
	var addressModel model.AddressBook
	if err := global.DB.Where("user_id=? and is_default= ?", value, 1).First(&addressModel).Error; err != nil {
		return nil, err
	}
	return &addressModel, nil
}

func (s AddressService) SetDefaultAddress(id int, c *gin.Context) {
	err := global.DB.Model(&model.AddressBook{}).Where("id = ?", id).Update("is_default", 1).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.EditError, c)
		return
	}

}

func (s AddressService) UpdateAddress(addReq req.AddressBookDTO, c *gin.Context) {
	var addressModel model.AddressBook
	if err := copier.Copy(&addressModel, &addReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ServerInternalError, c)
		return
	}
	if err := global.DB.Updates(&addressModel).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.SqlError, c)
		return
	}
}

func (s AddressService) AddAddress(addReq req.AddressBookDTO, c *gin.Context) {
	if addReq.UserId == 0 {
		value, exists := c.Get(enum.CurrentUserId)
		if !exists {
			global.Log.Error(code.UserError)
			res.FailWithMessage(code.UserError, c)
			return
		}
		userId, ok := value.(uint64)
		if !ok {
			global.Log.Error("Value is not of type int", userId)
			return
		}
		addReq.UserId = userId
	}

	if err := global.DB.Model(&model.AddressBook{}).Create(&addReq).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.SqlError, c)
		return
	}
}

func (s AddressService) DeleteAddress(id int, c *gin.Context) {
	if err := global.DB.Delete(&model.AddressBook{}, id).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.SqlError, c)
		return
	}

}
