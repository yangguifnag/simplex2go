package utils

import "gorm.io/gorm"

type Pagination struct {
	Total    int64 `json:"total"`
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
}

/* Init 初始化分页
total 总数
current 当前页
pageSize 每页数量
*/

func (p *Pagination) Init(total int64, current, pageSize int) {
	p.Total = total
	p.Current = current
	p.PageSize = pageSize
}

func PaginationGorm(page *Pagination) func(db *gorm.DB) *gorm.DB {
	// 如果当前页为-1，不分页 约定-1为不分页
	if page.Current == -1 {
		return nil
	}
	return func(db *gorm.DB) *gorm.DB {
		if page.PageSize == 0 {
			page.PageSize = 10
		}
		if page.Current == 0 {
			page.Current = 1
		}
		offset := (page.Current - 1) * page.PageSize
		return db.Offset(offset).Limit(page.PageSize)
	}

}
