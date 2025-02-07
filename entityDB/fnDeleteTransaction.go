package entityDB

import (
	"gorm.io/gorm"
	"log"
)

func (t *DbStruct[T, D]) DeleteByEntities2Transaction(tx *gorm.DB, m *[]T) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	result := tx.Table(t.GetTableName()).Delete(&m)
	if result.Error != nil {
		tx.Rollback()
		return result
	}
	return result
}

func (t *DbStruct[T, D]) DeleteByEntity2Transaction(tx *gorm.DB, m *T) *gorm.DB {
	return t.DeleteByEntities2Transaction(tx, &[]T{*m})
}

func (t *DbStruct[T, D]) DeleteByEntitiesWhitStatus2Transaction(tx *gorm.DB, m *[]T) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	result := tx.Table(t.GetTableName()).Model(&m).Updates(map[string]interface{}{"status": 0})
	if result.Error != nil {
		tx.Rollback()
		return result
	}
	return result

}

func (t *DbStruct[T, D]) DeleteByEntityWhitStatus2Transaction(tx *gorm.DB, m *T) *gorm.DB {
	return t.DeleteByEntitiesWhitStatus2Transaction(tx, &[]T{*m})
}
