package entityDB

import (
	"gorm.io/gorm"
	"log"
)

func (t *DbStruct[T, D]) InsertByEntities(m *[]T) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	db := t.GetGormDB()
	result := db.Table(t.GetTableName()).Create(&m)
	return result
}

func (t *DbStruct[T, D]) InsertByEntity(m *T) *gorm.DB {
	return t.InsertByEntities(&[]T{*m})
}
