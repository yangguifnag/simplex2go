package common

import "gorm.io/gorm/logger"

type MysqlConfigModule struct {
	User         string        `json:"user"`
	Pass         string        `json:"pass"`
	Port         string        `json:"port"`
	Host         string        `json:"host"`
	Name         string        `json:"name"`
	LoggerConfig logger.Config `json:"loggerConfig"`
}
