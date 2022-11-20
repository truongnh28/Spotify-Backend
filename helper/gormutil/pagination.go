package gormutil

import (
	"gorm.io/gorm"
)

const (
	minPageSize = 1
)

// Pagination defines pagination struct
type Pagination struct{ Size, Number int }

// Paginate page with limit and offset
func Paginate(paging Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset, limit := paging.calculateOffsetLimit()

		return db.Offset(offset).Limit(limit)
	}
}

func (p *Pagination) calculateOffsetLimit() (offset, limit int) {
	if p.Size < minPageSize {
		limit = -1
		offset = p.Number - 1
		return
	}

	if p.Number <= 0 {
		p.Number = 1
	}

	offset = (p.Number - 1) * p.Size
	limit = p.Size

	return
}
