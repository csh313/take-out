package userController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmshop/common/code"
	"hmshop/common/enum"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/service/userService"
	"strconv"
)

type AddressApi struct {
	service userService.AddressService
}

func (ad AddressApi) AddressList(c *gin.Context) {
	addressList := ad.service.AddressList(c)
	res.OkWithData(addressList, c)

}

func (ad AddressApi) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	address, err := ad.service.GetById(id, c)
	if err != nil {
		global.Log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.FailWithMessage("地址为空", c)
		} else {
			res.FailWithMessage(code.SqlError, c)

		}
		return
	}
	res.OkWithData(*address, c)

}

func (ad AddressApi) GetDefaultAddress(c *gin.Context) {
	value, exists := c.Get(enum.CurrentUserId)
	if !exists {
		global.Log.Error(errors.New("查询用户信息失败"))
		res.FailWithMessage(code.ReqError, c)
	}
	defaultAddress, err := ad.service.GetDefaultAddress(value, c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.OkWithMessage("您没有默认地址", c)
			return
		} else {
			global.Log.Error(err)
			res.FailWithMessage(code.SqlError, c)
			return
		}
	}
	res.OkWithData(defaultAddress, c)
}

func (ad AddressApi) SetDefaultAddress(c *gin.Context) {
	value := c.Query("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	ad.service.SetDefaultAddress(id, c)
	res.OkWithMessage(code.EditSuccess, c)
}

func (ad AddressApi) UpdateAddress(c *gin.Context) {
	var addReq req.AddressBookDTO
	if err := c.ShouldBindJSON(&addReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	ad.service.UpdateAddress(addReq, c)
	res.OkWithMessage(code.EditSuccess, c)
}

func (ad AddressApi) AddAddress(c *gin.Context) {
	var addreq req.AddressBookDTO
	if err := c.ShouldBindJSON(&addreq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	ad.service.AddAddress(addreq, c)
	res.OkWithMessage(code.AddSuccess, c)
}

func (ad AddressApi) DeleteAddress(c *gin.Context) {
	value := c.Query("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	ad.service.DeleteAddress(id, c)
	res.OkWithMessage(code.DeleteSuccess, c)
}
