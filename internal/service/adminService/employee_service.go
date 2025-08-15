package adminService

import (
	"errors"
	"hmshop/common/code"
	"hmshop/common/enum"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/model"
	pwd "hmshop/utils"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type EmployeeService struct {
}

func (EmployeeService) Login(c *gin.Context, req req.LoginDTO) string {
	var employeeModel model.Employee
	if err := global.DB.Where("username", req.UserName).First(&employeeModel).Error; err != nil {
		global.Log.Warn("用户名不存在")
	}
	//校验密码
	checkPwd := pwd.CheckPwd(employeeModel.Password, req.Password)
	if !checkPwd {
		global.Log.Debug("密码错误")
		res.FailWithMessage("密码错误", c)
	}
	//验证状态
	if employeeModel.Status == enum.Disable {
		res.FailWithMessage("用户被禁用", c)
	}
	//生成token
	token, err := pwd.GenerateToken(c, employeeModel.Id, global.AppConfig.Jwt.Admin.Name, global.AppConfig.Jwt.Admin.Secret)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}

	return token
}

func (EmployeeService) Register(c *gin.Context, req req.RegisterDTO) error {
	var employeeModel model.Employee
	if err := global.DB.Take(&employeeModel, "username", req.UserName).Error; err == nil {
		return errors.New("用户名已存在")
	}

	err := global.DB.Create(&model.Employee{
		Username: req.UserName,
		Password: pwd.HashPwd("123456"),
		IdNumber: req.IdNumber,
		Phone:    req.Phone,
		Sex:      req.Sex,
		Name:     req.Name,
	}).Error

	return err
}

func (EmployeeService) UpdateEmployee(c *gin.Context, req req.RegisterDTO, model model.Employee) error {
	uId, _ := c.Get(enum.CurrentId)

	if err := global.DB.Take(&model, uId).Error; err != nil {
		return err
	}
	maps := structs.Map(req)
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
	err := tx.Model(&model).Updates(DataMap).Error
	return err
}

func (EmployeeService) UpdatePassword(c *gin.Context, req req.EmployeeEditPassword, model model.Employee) error {
	uId, _ := c.Get(enum.CurrentId)
	if err := global.DB.Take(&model, uId).Error; err != nil {
		return errors.New("用户不存在")
	}

	checkPwd := pwd.CheckPwd(model.Password, req.OldPassword)
	if !checkPwd {
		return errors.New("原密码错误")
	}

	if err := global.DB.Model(&model).Update("password", pwd.HashPwd(req.NewPassword)).Error; err != nil {
		global.Log.Error(code.EditError)
		return err
	}
	return nil
}

func (EmployeeService) UpdateStatus(status int, id uint64, model model.Employee) error {
	if err := global.DB.Take(&model, id).Error; err != nil {
		return err
	}
	if err := global.DB.Model(&model).Update("status", status).Error; err != nil {
		global.Log.Error(code.EditError)
		return err
	}
	return nil

}
