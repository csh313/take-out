package provider

import (
	"hmshop/internal/api/adminController"
	"hmshop/internal/api/userController"

	"github.com/google/wire"
)

type AllController struct {
	//admin Controller
	EmployeeApi adminController.EmployeeApi
	CategoryApi adminController.CategoryApi
	CommonApi   adminController.CommonApi
	DishApi     adminController.DishApi
	OrderApi    adminController.OrderApi
	SetmealApi  adminController.SetmealApi
	ShopApi     adminController.ShopApi

	//user Controller
	AddressUApi  userController.AddressApi
	CategoryUApi userController.CategoryApi
	DishUApi     userController.DishApi
	OrderUApi    userController.OrderApi
	SetmealUApi  userController.SetmealApi
	ShopUApi     userController.ShopApi
	ShoppingUApi userController.ShoppingCartApi
	UserUApi     userController.UserApi
}

func NewController(employeeApi adminController.EmployeeApi,
	categoryApi adminController.CategoryApi,
	commonApi adminController.CommonApi,
	dishApi adminController.DishApi,
	orderApi adminController.OrderApi,
	setmealApi adminController.SetmealApi,
	shopApi adminController.ShopApi,

	addressUApi userController.AddressApi,
	categoryUApi userController.CategoryApi,
	dishUApi userController.DishApi,
	orderUApi userController.OrderApi,
	setmealUApi userController.SetmealApi,
	shopUApi userController.ShopApi,
	shoppingUApi userController.ShoppingCartApi,
	userUApi userController.UserApi) *AllController {
	return &AllController{
		EmployeeApi: employeeApi,
		CommonApi:   commonApi,
		DishApi:     dishApi,
		OrderApi:    orderApi,
		CategoryApi: categoryApi,
		SetmealApi:  setmealApi,
		ShopApi:     shopApi,

		AddressUApi:  addressUApi,
		CategoryUApi: categoryUApi,
		DishUApi:     dishUApi,
		OrderUApi:    orderUApi,
		SetmealUApi:  setmealUApi,
		ShopUApi:     shopUApi,
		ShoppingUApi: shoppingUApi,
		UserUApi:     userUApi,
	}
}

var ProviderController = wire.NewSet(
	NewController,
)
