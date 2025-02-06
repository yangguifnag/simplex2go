package utils

import "gorm.io/gorm"

type Pagination struct {
	Total    int64 `json:"total"`
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
}

type PaginationResult struct {
	Total    int64         `json:"total"`
	Current  int           `json:"current"`
	PageSize int           `json:"pageSize"`
	Row      []interface{} `json:"row"`
}

/*
	Init 初始化分页

total 总数
current 当前页
pageSize 每页数量
*/
func (p *Pagination) Init(total int64, current, pageSize int) {
	p.Total = total       // 总数
	p.Current = current   // 当前页
	p.PageSize = pageSize // 每页数量
}

// PaginationGorm 返回一个用于配置分页的Gorm处理器
// 此函数旨在与Gorm库一起使用，以实现数据库查询的分页功能
// 参数page是分页配置的指针，包含当前页和每页大小的信息
// 如果当前页设置为-1，则表示不分页，函数将返回nil
// 否则，函数将返回一个闭包，用于配置Gorm的Offset和Limit方法，以实现分页查询
func PaginationGorm(page *Pagination) func(db *gorm.DB) *gorm.DB {
	// 如果当前页为-1，不分页 约定-1为不分页
	if page.Current == -1 {
		return nil
	}
	// 返回一个闭包，用于配置分页
	return func(db *gorm.DB) *gorm.DB {
		// 如果每页大小未设置，则默认为10
		if page.PageSize == 0 {
			page.PageSize = 10
		}
		// 如果当前页未设置，则默认为第1页
		if page.Current == 0 {
			page.Current = 1
		}
		// 计算Offset值，即跳过前面多少条记录
		offset := (page.Current - 1) * page.PageSize
		// 使用Offset和Limit配置分页查询，并返回配置好的DB指针
		return db.Offset(offset).Limit(page.PageSize)
	}
}
