package entityDB

import (
	"gorm.io/gorm"
	"log"
)

func (t *DbStruct[T, D]) UpdateByEntity(m *T) *gorm.DB {

	return t.UpdateByEntities(&[]T{*m})
}

func (t *DbStruct[T, D]) UpdateByEntities(m *[]T) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Updates(&m)
	return db
}

func (t *DbStruct[T, D]) UpdateByEntitiesAndWhere(m *[]T, where string, params ...interface{}) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Where(where, params...).Updates(&m)
	return db
}

func (t *DbStruct[T, D]) UpdateByEntityAndWhere(m *T, where string, params ...interface{}) *gorm.DB {
	return t.UpdateByEntitiesAndWhere(&[]T{*m}, where, params...)
}

func (t *DbStruct[T, D]) UpdateByEntitiesAndWhereAndSelect(m *[]T, where string, params []interface{}, selector []interface{}) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Model(&m).Select("", selector...).Where(where, params...).Updates(&m)
	return db
}

func (t *DbStruct[T, D]) UpdateByEntityAndWhereAndSelect(m *T, where string, params []interface{}, selector []interface{}) *gorm.DB {
	return t.UpdateByEntitiesAndWhereAndSelect(&[]T{*m}, where, params, selector)
}

func (t *DbStruct[T, D]) UpdateByEntitiesAndWhereAndOmit(m *[]T, where string, params []interface{}, selector []string) *gorm.DB {
	if !t.check() {
		log.Println("DbStruct check failed")
		return nil
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Model(&m).Omit(selector...).Where(where, params...).Updates(&m)
	return db
}

func (t *DbStruct[T, D]) UpdateByEntityAndWhereAndOmit(m *T, where string, params []interface{}, selector []string) *gorm.DB {
	return t.UpdateByEntitiesAndWhereAndOmit(&[]T{*m}, where, params, selector)
}
