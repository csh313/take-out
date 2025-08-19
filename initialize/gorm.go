package initialize

import (
	"hmshop/config"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

func InitGorm(dsn string) *gorm.DB {

	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
		DontSupportRenameColumn:  true,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	//SlowQueryLog(db)
	//GormRateLimiter(db,rate.NewLimiter(500,1000))
	return db
}

func InitMysqlCluster(sql config.MysqlConf) *gorm.DB {
	mainDBStr := sql.Master.Dsn()
	slaveDBS := make([]gorm.Dialector, len(sql.Slave))
	for _, slaveDBStr := range sql.Slave {
		slaveDBS = append(slaveDBS, mysql.Open(slaveDBStr.Dsn()))
	}
	db, err := gorm.Open(mysql.Open(mainDBStr))
	if err != nil {
		panic("无法连接主库: " + err.Error())
	}
	// 配置读写分离插件
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{
			mysql.Open(mainDBStr),
		},
		Replicas: slaveDBS,
		Policy:   dbresolver.RoundRobinPolicy(), // 轮询策略
	}))
	if err != nil {
		panic("配置主从复制失败:%v" + err.Error())
	}
	return db
}
