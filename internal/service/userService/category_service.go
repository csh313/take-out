package userService

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/model"
)

type CategoryService struct{}

func (s CategoryService) GetCategory(c *gin.Context) []model.Category {
	var categoryModel []model.Category
	if err := global.DB.Find(&categoryModel).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
	}
	return categoryModel
}
