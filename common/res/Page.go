package res

import (
	"gorm.io/gorm"
	"hmshop/global"
	"hmshop/internal/api/req"
)

func PageList[T any](page req.PageInfo, model T) (list PageResult, err error) {
	list.Rows = make([]T, 0)

	offset := (page.Page - 1) * page.PageSize
	if offset < 0 {
		offset = 0
	}
	//如果pageSize为0，则查询所有
	if page.PageSize <= 0 {
		page.PageSize = -1
	}
	query := global.DB.Model(&model)
	if page.Name != "" {
		query = query.Where("name LIKE ?", "%"+page.Name+"%")
	}
	if page.Type != 0 {
		query = query.Where("type = ?", page.Type)
	}
	if page.Status != 0 {
		query = query.Where("status = ?", page.Status)
	}
	if page.CategoryId != 0 {
		query = query.Where("category_id = ?", page.CategoryId)
	}

	//global.DB.Where(&model).Find(&(list.Rows))
	query.Count(&(list.Total))

	err = query.Limit(page.PageSize).Offset(offset).Preload("SetMealDishes").Find(&list.Rows).Error

	return list, err
}

func PageListRow[T any](page req.PageInfo, model T) (*gorm.DB, int, int) {

	offset := (page.Page - 1) * page.PageSize
	if offset < 0 {
		offset = 0
	}
	//如果pageSize为0，则查询所有
	if page.PageSize <= 0 {
		page.PageSize = -1
	}
	query := global.DB.Model(&model)
	if page.Name != "" {
		query = query.Where("name LIKE ?", "%"+page.Name+"%")
	}
	if page.Type != 0 {
		query = query.Where("type = ?", page.Type)
	}
	if page.Status != 0 {
		query = query.Where("status = ?", page.Status)
	}
	if page.CategoryId != 0 {
		query = query.Where("category_id = ?", page.CategoryId)
	}

	//global.DB.Where(&model).Find(&(list.Rows))
	//query.Count(&(list.Total))

	//query = query.Limit(page.PageSize).Offset(offset)

	return query, page.PageSize, offset
}
