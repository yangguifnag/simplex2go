package entityDB

import (
	"github.com/yangguifnag/simplex2go/utils"
	"log"
)

func (t *DbStruct[T, D]) SelectAll2Page(page *utils.Pagination) *utils.PaginationResult[T] {
	return t.SelectByWhere2Page(page, "1=?", "1")
}

func (t *DbStruct[T, D]) SelectByEntity2Page(page *utils.Pagination, m *T) *utils.PaginationResult[T] {
	var result []T
	if !t.check() {
		log.Println("DbStruct check failed")
		return &utils.PaginationResult[T]{}
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Where(&m).Scopes(utils.PaginationGorm(page)).Find(&result)
	return &utils.PaginationResult[T]{
		Row:      result,
		Total:    page.Total,
		PageSize: page.PageSize,
		Current:  page.Current,
	}
}

func (t *DbStruct[T, D]) SelectOne2Page(page *utils.Pagination) *utils.PaginationResult[T] {
	return t.SelectAll2Page(page)
}

func (t *DbStruct[T, D]) SelectByWhere2Page(page *utils.Pagination, where string, params ...interface{}) *utils.PaginationResult[T] {
	var result []T
	if !t.check() {
		log.Println("DbStruct check failed")
		return &utils.PaginationResult[T]{}
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Where(where, params...).Scopes(utils.PaginationGorm(page)).Find(&result)
	return &utils.PaginationResult[T]{
		Row:      result,
		Total:    page.Total,
		PageSize: page.PageSize,
		Current:  page.Current,
	}
}

func (t *DbStruct[T, D]) SelectByWhereWhitStatus2Page(page *utils.Pagination, where string, params ...interface{}) *utils.PaginationResult[T] {
	var result []T
	if !t.check() {
		log.Println("DbStruct check failed")
		return &utils.PaginationResult[T]{}
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Where(where, params...).Where(`status = 1`).Scopes(utils.PaginationGorm(page)).First(&result)
	return &utils.PaginationResult[T]{
		Row:      result,
		Total:    page.Total,
		PageSize: page.PageSize,
		Current:  page.Current,
	}
}

func (t *DbStruct[T, D]) SelectAllWithStatus2Page(page *utils.Pagination) *utils.PaginationResult[T] {
	return t.SelectByWhere2Page(page, "status=?", 1)
}

func (t *DbStruct[T, D]) SelectByEntityWithStatus2Page(page *utils.Pagination, m *T) *utils.PaginationResult[T] {
	var result []T
	if !t.check() {
		log.Println("DbStruct check failed")
		return &utils.PaginationResult[T]{}
	}
	db := t.GetGormDB()
	db.Table(t.GetTableName()).Where(&m).Where(`status = 1`).Find(&result)
	return &utils.PaginationResult[T]{
		Row:      result,
		Total:    page.Total,
		PageSize: page.PageSize,
		Current:  page.Current,
	}
}
