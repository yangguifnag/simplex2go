package entityDB

import (
	"gorm.io/gorm"
	"log"
)

func (t *DbStruct[T, D]) UpdateByEntity2Transaction(tx *gorm.DB, m *T) *gorm.DB {

	return t.UpdateByEntities2Transaction(tx, &[]T{*m})
}

func (t *DbStruct[T, D]) UpdateByEntities2Transaction(tx *gorm.DB, m *[]T) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	result := tx.Table(t.GetTableName()).Updates(&m)
	if result.Error != nil {
		tx.Rollback()
		return result
	}
	return result
}

func (t *DbStruct[T, D]) UpdateByEntitiesAndWhere2Transaction(tx *gorm.DB, m *[]T, where string, params ...interface{}) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	result := tx.Table(t.GetTableName()).Where(where, params...).Updates(&m)
	if result.Error != nil {
		tx.Rollback()
		return result
	}
	return result
}

func (t *DbStruct[T, D]) UpdateByEntityAndWhere2Transaction(tx *gorm.DB, m *T, where string, params ...interface{}) *gorm.DB {
	return t.UpdateByEntitiesAndWhere2Transaction(tx, &[]T{*m}, where, params...)
}

func (t *DbStruct[T, D]) UpdateByEntitiesAndWhereAndSelect2Transaction(tx *gorm.DB, m *[]T, where string, params []interface{}, selector []interface{}) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	result := tx.Table(t.GetTableName()).Model(&m).Select("", selector...).Where(where, params...).Updates(&m)
	if result.Error != nil {
		tx.Rollback()
		return result
	}
	return result
}

func (t *DbStruct[T, D]) UpdateByEntityAndWhereAndSelect2Transaction(tx *gorm.DB, m *T, where string, params []interface{}, selector []interface{}) *gorm.DB {
	return t.UpdateByEntitiesAndWhereAndSelect2Transaction(tx, &[]T{*m}, where, params, selector)
}

func (t *DbStruct[T, D]) UpdateByEntitiesAndWhereAndOmit2Transaction(tx *gorm.DB, m *[]T, where string, params []interface{}, selector []string) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	result := tx.Table(t.GetTableName()).Model(&m).Omit(selector...).Where(where, params...).Updates(&m)
	if result.Error != nil {
		tx.Rollback()
		return result
	}
	return result
}

func (t *DbStruct[T, D]) UpdateByEntityAndWhereAndOmit2Transaction(tx *gorm.DB, m *T, where string, params []interface{}, selector []string) *gorm.DB {
	return t.UpdateByEntitiesAndWhereAndOmit2Transaction(tx, &[]T{*m}, where, params, selector)
}
