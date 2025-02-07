package entityDB

import (
	"gorm.io/gorm"
	"log"
)

type TransactionAction func(tx *gorm.DB) *gorm.DB
type TransactionActionV2 func(tx *gorm.DB, tranFn ...any) *gorm.DB

type TransactionActionV2Struct struct {
	Method TransactionActionV2
	Args   []any
}

type DbStruct[T, D any] struct {
	Entity  *T `json:"entity"`
	MysqlBD D
}

func (t *DbStruct[T, D]) check() bool {
	flag := true
	if gormDB, ok := any(t.MysqlBD).(*gorm.DB); ok {
		flag = gormDB != nil
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
	if entity, ok := any(t.Entity).(interface {
		GetTableName() string
	}); ok {
		return entity.GetTableName()
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

func (t *DbStruct[T, D]) Transaction(txFn TransactionAction) *gorm.DB {

	tx := t.TransactionBegin()

	result := txFn(tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := result.Error; err != nil {
		log.Println("Transaction error:", err)
		return result
	}

	return t.TransactionCommit(result)
}

func (t *DbStruct[T, D]) TransactionV2(txFns []TransactionActionV2Struct) *gorm.DB {

	tx := t.TransactionBegin()
	result := tx
	for _, txFn := range txFns {
		result = txFn.Method(result, txFn.Args...)
		if err := result.Error; err != nil {
			log.Println("Transaction error:", err)
			tx.Rollback()
			return tx
		}
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return t.TransactionCommit(result)
}
