package adminService

import (
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmshop/common/code"
	"hmshop/common/enum"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/api/resp"
	"hmshop/internal/model"
	"time"
)

type SetmealService struct {
}

func (s SetmealService) AddSetmeal(req req.SetMealDTO, c *gin.Context) error {
	//tx := global.DB.WithContext(c).Begin()
	tx := global.DB.WithContext(c)

	//查找是否已包含该套餐
	if err := tx.Take(&model.SetMeal{}, "name", req.Name).RowsAffected; err != 0 {
		return errors.New("套餐已存在")
	}

	tx = tx.Begin()
	var setmealModel = model.SetMeal{
		CategoryId:    req.CategoryId,
		Name:          req.Name,
		Price:         req.Price,
		Status:        req.Status,
		Description:   req.Description,
		Image:         req.Image,
		SetMealDishes: req.SetMealDishs,
	}
	if err := tx.Create(&setmealModel).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return errors.New(code.AddError + "套餐")
	}
	//方法二：尝试不同gorm的特性，分别在两个表中添加数据
	//var setmealDishModel = req.SetMealDishs
	//for i, _ := range setmealDishModel {
	//	setmealDishModel[i].SetmealId = setmealModel.Id
	//}
	//fmt.Println(setmealDishModel)
	//if err := tx.Create(&setmealDishModel).Error; err != nil {
	//	global.Log.Error(err.Error())
	//	tx.Rollback()
	//	return errors.New(code.AddError + "套餐菜品")
	//}
	tx.Commit()
	//抛出异常
	defer func() {
		if r := recover(); r != nil {
			//遇到错误时回滚
			tx.Rollback()
		}
	}()
	return nil
}

func (s SetmealService) GetById(id int, c *gin.Context) (*resp.SetMealWithDishByIdVo, error) {
	var setmealModel = model.SetMeal{}
	if err := global.DB.Preload("SetMealDishes").Take(&setmealModel, "id=?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.FailWithMessage(code.DataNotFound, c)
			return nil, err
		}
		global.Log.Error(err.Error())
		return nil, errors.New(code.QueryError)
	}
	// 查询 CategoryName（根据 CategoryId 获取分类名称）
	var categoryName string
	if err := global.DB.Model(&model.Category{}).Where("id=?", setmealModel.CategoryId).
		Select("name").Take(&categoryName).Error; err != nil {
		global.Log.Error(err.Error())
		return nil, errors.New(code.QueryError)
	}
	result := &resp.SetMealWithDishByIdVo{
		Id:            setmealModel.Id,
		CategoryId:    setmealModel.CategoryId,
		CategoryName:  categoryName,
		Description:   setmealModel.Description,
		Image:         setmealModel.Image,
		Name:          setmealModel.Name,
		Price:         setmealModel.Price,
		SetmealDishes: setmealModel.SetMealDishes, // 直接映射
		Status:        setmealModel.Status,
		UpdateTime:    setmealModel.UpdateTime,
	}
	return result, nil
}

func (s SetmealService) PageQuery(page req.PageInfo, c *gin.Context) (res.PageResult, error) {
	list, err := res.PageList(page, model.SetMeal{})
	return list, err
}

func (s SetmealService) DeleteByIds(list []string, c *gin.Context) error {
	tx := global.DB.WithContext(c).Begin()
	tx.Begin()
	if err := tx.Where("setmeal_id in (?)", list).Delete(&model.SetMealDish{}).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return errors.New(code.DeleteError)
	}
	if err := tx.Where("id in (?)", list).Delete(&model.SetMeal{}).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return errors.New(code.DeleteError)
	}

	tx.Commit()

	//抛出异常
	defer func() {
		if r := recover(); r != nil {
			//遇到错误时回滚
			tx.Rollback()
		}
	}()

	return nil
}

func (s SetmealService) UpdateMeal(mealReq req.SetMealDTO, c *gin.Context) error {
	var setmealModel = model.SetMeal{}
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

	//先删除菜品与套餐的关联
	if err := tx.Where("setmeal_id=?", mealReq.Id).Delete(&model.SetMealDish{}).Error; err != nil {
		global.Log.Error(err.Error(), "菜品与套餐")
		tx.Rollback()
		return errors.New(code.DeleteError)
	}

	if err := tx.Preload("SetMealDishes").Take(&setmealModel, "id=?", mealReq.Id).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return errors.New(code.ReqError)
	}

	maps := structs.Map(mealReq)
	var DataMap = map[string]any{}
	delete(maps, "SetMealDishs")
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

	fmt.Println(DataMap, "=========")
	// 手动设置更新时间和更新人
	DataMap["UpdateTime"] = time.Now()
	DataMap["UpdateUser"] = uId
	if err := tx.Model(&setmealModel).Updates(DataMap).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return errors.New(code.EditError)
	}

	//添加setmealdish数据
	setmealDishesReq := mealReq.SetMealDishs
	for i, _ := range setmealDishesReq {
		setmealDishesReq[i].SetmealId = mealReq.Id

	}
	fmt.Println(setmealDishesReq)
	if err := tx.Create(&setmealDishesReq).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return errors.New(code.EditError)
	}

	tx.Commit()
	return nil
}

func (s SetmealService) SetStatus(id uint64, status int, c *gin.Context) error {

	var model model.SetMeal
	Rows := global.DB.Preload("SetMealDishes").Take(&model, id).RowsAffected
	if Rows <= 0 {
		global.Log.Error("该套餐不存在")
		return errors.New(code.DataNotFound)
	}
	if err := global.DB.WithContext(c).Model(&model).Update("status", status).Error; err != nil {
		global.Log.Error(err.Error())
		return errors.New(code.EditError)
	}
	return nil
}
