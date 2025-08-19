package userService

import (
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/resp"
	"hmshop/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SetmealService struct {
}

func (s SetmealService) SetmealByCategoryId(id int, c *gin.Context) []resp.SetMealPageQueryVo {
	var setmealModel []model.SetMeal
	if err := global.DBs.Where("category_id = ?", id).Find(&setmealModel).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
	}
	var setMealPageQueryVos []resp.SetMealPageQueryVo
	if err := copier.Copy(&setMealPageQueryVos, &setmealModel); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
	}
	return setMealPageQueryVos

}

func (s SetmealService) GetDishById(id int, c *gin.Context) ([]resp.DishItemVO, error) {
	var dishItems []resp.DishItemVO
	err := global.DBs.Table("setmeal_dish").
		Select("setmeal_dish.name,setmeal_dish.copies,dish.image,dish.description").
		Joins("left join dish on setmeal_dish.dish_id=dish.id").
		Where("setmeal_dish.setmeal_id=?", id).Scan(&dishItems).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
		return nil, err
	}
	return dishItems, nil
}
