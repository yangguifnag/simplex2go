package db

import (
	"database/sql"
	"github.com/yangguifnag/simplex2go/common"
	"github.com/yangguifnag/simplex2go/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)
import "gorm.io/gorm"

type MysqlDB struct {
	Config map[string]common.MysqlConfigModule
	db     map[string]*gorm.DB
}

func (db *MysqlDB) GetDB(name string) *gorm.DB {
	if db.db == nil {
		db.db = make(map[string]*gorm.DB)
	}

	if db.db[name] != nil {
		return db.db[name]
	}

	module := db.Config[name]
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Nanosecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,     // Log level
			Colorful:      true,            // 禁用彩色打印

		},
	)
	DB, _ := gorm.Open(mysql.Open(utils.GetDSN(module)), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	sqlDB, _ := DB.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatalf("关闭数据库连接失败: %v", err)
		}
	}(sqlDB)
	sqlDB.SetMaxIdleConns(100) //设置最大连接数
	sqlDB.SetMaxOpenConns(100) //设置最大的空闲连接数
	return DB
}
