package initialize

import (
	"hmshop/config"
	"hmshop/global"
	"hmshop/internal/routers"
	"hmshop/internal/service/kafkaService"
	"hmshop/logger"
	util "hmshop/utils"

	"github.com/gin-gonic/gin"
)

func GlobalInit() *gin.Engine {

	//配置文件初始化
	global.AppConfig = config.InitConfig()
	global.Log = logger.NewLogger(global.AppConfig.Log.Level, global.AppConfig.Log.FilePath)
	global.DB = InitGorm(global.AppConfig.Datasource.Dsn())
	global.DBs = InitMysqlCluster(global.AppConfig.MysqlConf)
	global.Redis = InitRedis()
	global.KafkaConfig = InitKafkaConfig()
	util.InitAliOss()
	global.Producer = InitProducer()
	global.Consumer = InitConsumer()
	go kafkaService.ConsumeMessages(global.Consumer)

	router := routers.RouterInit()
	return router
}
