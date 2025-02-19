package userService

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/resp"
	"hmshop/internal/model"
)

type DishService struct{}

func (DishService) ListDish(categoryId int, c *gin.Context) []resp.DishVo {
	var dishModel []model.Dish
	if err := global.DB.Preload("Flavors").Where("category_id=?", categoryId).Find(&dishModel).Error; err != nil {
		res.FailWithMessage(code.QueryError, c)
	}
	var dishVoList []resp.DishVo
	err := copier.Copy(&dishVoList, &dishModel)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
		return nil
	}
	return dishVoList
}
