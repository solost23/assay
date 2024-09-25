package initialize

import (
	"assay/infra/global"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DefaultStringSize = 100
)

func initMysql() {
	mysqlConfig := global.ServerConfig.MySQL

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               mysqlConfig.DataSourceName,
		DefaultStringSize: DefaultStringSize,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		zap.S().Panic("failed to connect database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Panic("failed to get db connection: ", err)
	}
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConfig.MaxConnLifeTime) * time.Second)

	global.DB = db
}
