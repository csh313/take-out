package userService

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hmshop/common/code"
	"hmshop/common/enum"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/model"
	"hmshop/internal/service/adminService"
)

type ShoppingCartService struct {
}

func (sc ShoppingCartService) List(value any, c *gin.Context) []model.ShoppingCart {
	var shoppingCart []model.ShoppingCart
	if err := global.DB.Find(&shoppingCart).Where("user_id=?", value).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.SqlError, c)
	}
	return shoppingCart
}

func (sc ShoppingCartService) AddCart(cartReq req.ShoppingCartDTO, c *gin.Context) error {
	var shoppingCart model.ShoppingCart
	err := copier.Copy(&shoppingCart, &cartReq)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ServerInternalError, c)
		return err
	}

	user_id, exists := c.Get(enum.CurrentUserId)
	if !exists {
		global.Log.Error(errors.New(code.UserError))
		res.FailWithMessage(code.UserError, c)
		return err
	}
	userId := user_id.(uint64)
	shoppingCart.UserId = userId
	//查询当前商品是否在购物车中
	shoppingCartList, err := sc.QueryShoppingCart(shoppingCart, c)
	if err != nil {
		return err
	}

	if shoppingCartList != nil && len(shoppingCartList) == 1 {
		//购物车中存在该商品
		shoppingCart = shoppingCartList[0]
		shoppingCart.Number++
		if err := global.DB.Updates(&shoppingCart).Error; err != nil {
			global.Log.Error(err)
			res.FailWithMessage(code.SqlError, c)
			return err
		}
	} else {
		//如果不存在就插入数据，数量为1
		//如果是菜品，就添加菜品信息
		if cartReq.DishId != 0 {
			dish, err := adminService.DishService{}.GetById(c, cartReq.DishId)
			if err != nil {
				return err
			}
			shoppingCart.Image = dish.Image
			shoppingCart.Name = dish.Name
			shoppingCart.Amount = dish.Price
		} else {
			//添加的是套餐
			setmeal, err := adminService.SetmealService{}.GetById(cartReq.SetmealId, c)
			if err != nil {
				return err
			}
			shoppingCart.Image = setmeal.Image
			shoppingCart.Name = setmeal.Name
			shoppingCart.Amount = setmeal.Price
		}
		shoppingCart.Number = 1

		if err := global.DB.Create(&shoppingCart).Error; err != nil {
			global.Log.Error(err)
			res.FailWithMessage(code.SqlError, c)
			return err
		}
	}

	return nil

}

func (sc ShoppingCartService) Delete(cartReq req.ShoppingCartDTO, c *gin.Context) error {
	var shoppingCart model.ShoppingCart
	if err := copier.Copy(&shoppingCart, &cartReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ServerInternalError, c)
		return err
	}
	value, exists := c.Get(enum.CurrentUserId)
	if !exists {
		global.Log.Error(errors.New(code.UserError))
		res.FailWithMessage(code.UserError, c)
		return errors.New("查询用户信息失败")
	}
	userId := value.(uint64)
	shoppingCart.UserId = userId
	shoppingCartList, err := sc.QueryShoppingCart(shoppingCart, c)
	if err != nil {
		return err
	}
	if shoppingCartList != nil && len(shoppingCartList) != 0 {
		shoppingCart = shoppingCartList[0]
		if shoppingCart.Number == 1 {
			//只有一份就直接删除
			if err := global.DB.Delete(&shoppingCart).Error; err != nil {
				global.Log.Error(err)
				res.FailWithMessage(code.SqlError, c)
				return err
			}
		} else {
			shoppingCart.Number--
			if err := global.DB.Updates(&shoppingCart).Error; err != nil {
				global.Log.Error(err)
				res.FailWithMessage(code.SqlError, c)
				return err
			}
		}
	}
	return nil

}

func (sc ShoppingCartService) QueryShoppingCart(shoppingCart model.ShoppingCart, c *gin.Context) ([]model.ShoppingCart, error) {
	var CartList []model.ShoppingCart
	if err := global.DB.Find(&CartList, shoppingCart).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.SqlError, c)
		return nil, err
	}
	return CartList, nil
}

func (sc ShoppingCartService) Clean(c *gin.Context) error {
	value, exists := c.Get(enum.CurrentUserId)
	if !exists {
		global.Log.Error(errors.New(code.UserError))
		res.FailWithMessage(code.UserError, c)
		return errors.New(code.UserError)
	}
	userId := value.(uint64)
	if err := global.DB.Where("user_id=?", userId).Delete(&model.ShoppingCart{}).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.SqlError, c)
		return err
	}
	return nil
}
