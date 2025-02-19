package main

import (
	"hmshop/global"
	"hmshop/initialize"
)

func main() {
	//初始化配置
	router := initialize.GlobalInit()
	router.Run(":" + global.AppConfig.Server.Port)
	//n := 1
	//var num *int = &n
	//fmt.Println(*num)
	//fmt.Println(num)
}
