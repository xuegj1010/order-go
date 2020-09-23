package datasource

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"order-go/common"
	"order-go/models"
)

func NewMysqlEngine() *gorm.DB {
	cfg := common.LoadCfg()
	dsn := cfg.MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // 禁用彩色打印
			},
		),
	})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
