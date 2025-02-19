package adminService

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/enum"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/model"
	"time"
)

type CategoryService struct {
}

func (cs CategoryService) AddCategory(c *gin.Context, request req.CategoryDTO) error {

	var categoryModel model.Category
	tx := global.DB.WithContext(c)
	if err := tx.Take(&categoryModel, "name", request.Name).Error; err != nil {
		global.Log.Warn("该分类已存在")
		return err
	}
	err := tx.Create(&model.Category{
		Type: request.Type,
		Name: request.Name,
		Sort: request.Sort,
	}).Error
	return err
}

func (cs CategoryService) List(c *gin.Context, typeId int) ([]model.Category, error) {
	var categoryModel []model.Category
	err := global.DB.WithContext(c).Find(&categoryModel, "type = ?", typeId).Error

	return categoryModel, err
}

func (cs CategoryService) DeleteById(c *gin.Context, id uint64) error {
	var model model.Category
	if err := global.DB.Take(&model, "id = ?", id).Error; err != nil {
		global.Log.Error(code.DataNotFound)
		return err
	}
	err := global.DB.Delete(&model).Error
	return err

}

func (cs CategoryService) EditCategory(c *gin.Context, request req.CategoryDTO) error {
	var CategoryModel model.Category
	uId, _ := c.Get(enum.CurrentId)
	if err := global.DB.Take(&CategoryModel, "id = ?", request.Id).Error; err != nil {
		global.Log.Error(code.DataNotFound)
		return err
	}
	maps := structs.Map(request)
	var DataMap = map[string]any{}
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}
	// 手动设置更新时间和更新人
	DataMap["UpdateTime"] = time.Now()
	DataMap["UpdateUser"] = uId
	tx := global.DB.WithContext(c) // 将 gin.Context 上下文传递给 gorm
	//fmt.Println(model, "----model")
	err := tx.Model(&CategoryModel).Updates(DataMap).Error
	return err
}

func (cs CategoryService) SetStatus(c *gin.Context, id uint64, status int) error {
	var CategoryModel model.Category

	if err := global.DB.Take(&CategoryModel, "id = ?", id).Error; err != nil {
		global.Log.Error(code.DataNotFound)
		return err
	}
	err := global.DB.WithContext(c).Model(&CategoryModel).Update("status", status).Error
	return err
}
