package userService

import (
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/model"

	"github.com/gin-gonic/gin"
)

type CategoryService struct{}

func (s CategoryService) GetCategory(c *gin.Context) []model.Category {
	var categoryModel []model.Category
	if err := global.DBs.Find(&categoryModel).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
	}
	return categoryModel
}
