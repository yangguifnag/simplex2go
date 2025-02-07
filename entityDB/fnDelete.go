package entityDB

import (
	"gorm.io/gorm"
	"log"
)

func (t *DbStruct[T, D]) DeleteByEntities(m *[]T) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	db := t.GetGormDB()
	result := db.Table(t.GetTableName()).Delete(&m)
	return result
}

func (t *DbStruct[T, D]) DeleteByEntity(m *T) *gorm.DB {
	return t.DeleteByEntities(&[]T{*m})
}

func (t *DbStruct[T, D]) DeleteByEntitiesWhitStatus(m *[]T) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Model(&m).Updates(map[string]interface{}{"status": 0})
	return db

}

func (t *DbStruct[T, D]) DeleteByEntityWhitStatus(m *T) *gorm.DB {
	return t.DeleteByEntitiesWhitStatus(&[]T{*m})
}
