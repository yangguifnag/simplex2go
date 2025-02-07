package entityDB

import (
	"gorm.io/gorm"
	"log"
)

func (t *DbStruct[T, D]) InsertByEntities2Transaction(tx *gorm.DB, m *[]T) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	result := tx.Table(t.GetTableName()).Create(&m)
	if result.Error != nil {
		tx.Rollback()
		return result
	}
	return result
}

func (t *DbStruct[T, D]) InsertByEntity2Transaction(tx *gorm.DB, m *T) *gorm.DB {
	return t.InsertByEntities2Transaction(tx, &[]T{*m})
}
