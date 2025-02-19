package adminService

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/enum"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/api/resp"
	"hmshop/internal/model"
	"time"
)

type DishService struct {
}

func (DishService) AddDish(c *gin.Context, dishReq req.DishDTO) error {

	//开启事务
	tx := global.DB.Begin()
	//抛出异常
	defer func() {
		if r := recover(); r != nil {
			//遇到错误时回滚
			tx.Rollback()
		}
	}()
	//在事务中执行一些数据库操作（此时使用 'tx'，而不是 'db'）
	var dishModel = model.Dish{
		Id:          dishReq.Id,
		Name:        dishReq.Name,
		CategoryId:  dishReq.CategoryId,
		Price:       dishReq.Price,
		Image:       dishReq.Image,
		Description: dishReq.Description,
		Status:      dishReq.Status,
		Flavors:     dishReq.Flavors,
	}
	if err := tx.WithContext(c).Create(&dishModel).Error; err != nil {
		global.Log.Error(err.Error(), "创建菜品失败")
		tx.Rollback()
		return err
	}

	//执行
	tx.Commit()

	return nil
}

func (s DishService) DishPage(c *gin.Context, dishReq req.PageInfo) (res.PageResult, error) {
	var list res.PageResult
	//list.Rows = make([]any, 0)
	list.Rows = make([]resp.DishPageVo, 0)
	var dishModel model.Dish
	query, pageSize, offset := res.PageListRow(dishReq, dishModel)
	query = query.Limit(pageSize).Offset(offset)

	query.Count(&(list.Total))
	err := query.Select("dish.id, dish.name, dish.category_id, dish.price, dish.image, dish.description, dish.status, dish.update_time, category.name as category_name").
		Joins("left join category on category.id = dish.category_id").Scan(&list.Rows).Error

	return list, err
}

func (s DishService) GetById(c *gin.Context, id int) (model.Dish, error) {
	var dishModel = model.Dish{}
	err := global.DB.Preload("Flavors").Where("id = ?", id).First(&dishModel).Error
	if err != nil {
		global.Log.Error(err.Error())
	}
	return dishModel, err
}

func (s DishService) List(categoryId uint64, c *gin.Context) ([]model.Dish, error) {
	var list []model.Dish
	err := global.DB.Preload("Flavors").Where("category_id =?", categoryId).Find(&list).Error
	return list, err

}

func (s DishService) DeleteDish(list []string, c *gin.Context) error {
	err := global.DB.Where("id in (?)", list).Delete(&model.Dish{}).Error
	return err
}
func (s DishService) UpdateDish(c *gin.Context, dishReq req.DishUpdateDTO) error {
	var dishModel = model.Dish{}
	uId, _ := c.Get(enum.CurrentId)
	//开启事务
	tx := global.DB.Begin()
	//抛出异常
	defer func() {
		if r := recover(); r != nil {
			//遇到错误时回滚
			tx.Rollback()
		}
	}()
	if err := tx.Where("dish_id=?", dishReq.Id).Delete(&model.DishFlavor{}).Error; err != nil {
		global.Log.Error(err.Error(), "修改菜品口味失败")
		tx.Rollback()
		return err
	}
	if err := tx.Preload("Flavors").Take(&dishModel, "id = ?", dishReq.Id).Error; err != nil {
		global.Log.Error(code.DataNotFound)
		tx.Rollback()
		return err
	}
	//删除全部口味，将传进来的再添加上去
	maps := structs.Map(dishReq)
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
		case float64:
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
	if err := tx.Model(&dishModel).Updates(DataMap).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return err
	}
	dishFlavor := dishReq.Flavors
	if err := tx.Create(&dishFlavor).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return err
	}
	//执行
	tx.Commit()
	return nil
}
