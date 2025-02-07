package entityDB

import (
	"gorm.io/gorm"
	"log"
)

type DbStruct[T, D any] struct {
	Entity  T `json:"entity"`
	MysqlBD D
}

func (t *DbStruct[T, D]) check() bool {
	flag := t.MysqlBD != nil && t.Entity != nil
	if gormDB, ok := any(t.MysqlBD).(*gorm.DB); ok {
		flag = flag && gormDB != nil
	}
	return flag
}

func (t *DbStruct[T, D]) GetGormDB() *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	if gormDB, ok := any(t.MysqlBD).(*gorm.DB); ok {
		return gormDB
	}
	return nil
}

func (t *DbStruct[T, D]) GetTableName() string {
	if !t.check() {
		log.Println("DbStruct check failed")
		return ""
	}
	if tableName, ok := any(t.Entity).(interface {
		GetTableName() string
	}); ok {
		return tableName.GetTableName()
	}
	return ""
}

func (t *DbStruct[T, D]) TransactionBegin() *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	db := t.GetGormDB()
	tx := db.Begin()
	return tx
}

func (t *DbStruct[T, D]) TransactionCommit(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		tx.Commit()
	}
	return tx
}
