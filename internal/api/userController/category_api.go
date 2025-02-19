package userController

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/internal/service/userService"
)

type CategoryApi struct {
	service userService.CategoryService
}

func (a CategoryApi) GetCategoryByType(c *gin.Context) {
	categoryModel := a.service.GetCategory(c)
	res.Ok(categoryModel, code.QuerySuccess, c)
}
