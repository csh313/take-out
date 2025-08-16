package adminController

import (
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/model"
	"hmshop/internal/service/adminService"
	pwd "hmshop/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeApi struct {
	service adminService.EmployeeService
}

func (em EmployeeApi) Login(c *gin.Context) {
	employeeLogin := req.LoginDTO{}
	if err := c.ShouldBind(&employeeLogin); err != nil {
		global.Log.Debug(code.ReqError)
		return
	}
	resq := em.service.Login(c, employeeLogin)

	res.OkWithData(resq, c)
}

func (em EmployeeApi) Register(c *gin.Context) {
	resq := req.RegisterDTO{}
	if err := c.ShouldBind(&resq); err != nil {
		global.Log.Debug(code.ReqError)
		return
	}
	err := em.service.Register(c, resq)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("注册成功", c)
}

func (em EmployeeApi) Logout(c *gin.Context) {
	token := c.GetHeader(global.AppConfig.Jwt.Admin.Name)
	err := pwd.LogoutToken(c, token)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("注销成功", c)
}

func (em EmployeeApi) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	var employeeModel model.Employee
	if err := global.DBs.Where("id", id).First(&employeeModel).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(employeeModel, c)
}

func (em EmployeeApi) PageEmployee(c *gin.Context) {
	var page req.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		global.Log.Debug(code.ReqError)
		res.FailWithMessage(err.Error(), c)
		return
	}
	var employeeModel model.Employee
	list, err := res.PageList(page, employeeModel)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(list, c)
}

func (em EmployeeApi) UpdateEmployee(c *gin.Context) {
	var req req.RegisterDTO
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Debug(code.ReqError)
		res.FailWithMessage(err.Error(), c)
		return
	}
	var employeeModel model.Employee
	err := em.service.UpdateEmployee(c, req, employeeModel)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("修改成功", c)
}

func (em EmployeeApi) UpdatePassword(c *gin.Context) {
	var req req.EmployeeEditPassword
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Debug(code.ReqError)
		res.FailWithMessage(err.Error(), c)
		return
	}
	var employeeModel model.Employee
	err := em.service.UpdatePassword(c, req, employeeModel)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(code.EditSuccess, c)
}

func (em EmployeeApi) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		global.Log.Debug(code.ReqError)
		res.FailWithMessage(err.Error(), c)
		return
	}
	err = em.service.UpdateStatus(status, id, model.Employee{})
	if err != nil {
		global.Log.Debug(code.EditError)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(code.EditSuccess, c)

}
