package entityDB

import "log"

func (t *DbStruct[T, D]) SelectAll() []T {
	var result []T
	result = t.SelectByWhere("1 = ?", 1)
	return result
}

func (t *DbStruct[T, D]) SelectByEntity(m *T) []T {
	var result []T
	if !t.check() {
		log.Println("DbStruct check failed")
		return make([]T, 0)
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Where(&m).Find(&result)
	return result
}

func (t *DbStruct[T, D]) SelectOne() T {
	var result T
	result = t.SelectAll()[0]
	return result
}

func (t *DbStruct[T, D]) SelectByWhere(where string, params ...interface{}) []T {
	var result []T
	if !t.check() {
		log.Println("DbStruct check failed")
		return make([]T, 0)
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Where(where, params...).Find(&result)
	return result
}

func (t *DbStruct[T, D]) SelectByWhereWhitStatus(where string, params ...interface{}) []T {
	var result []T
	if !t.check() {
		log.Println("DbStruct check failed")
		return make([]T, 0)
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Where(where, params...).Where(`status = 1`).First(&result)
	return result
}

func (t *DbStruct[T, D]) SelectAllWithStatus() []T {
	var result []T
	result = t.SelectByWhere("status=?", 1)
	return result
}

func (t *DbStruct[T, D]) SelectByEntityWithStatus(P T) []T {
	var result []T
	if !t.check() {
		log.Println("DbStruct check failed")
		return make([]T, 0)
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Where(P).Where(`status = 1`).Find(&result)
	return result
}
