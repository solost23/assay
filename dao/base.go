package dao

import "gorm.io/gorm"

const (
	MaxSize     = 100
	BatchSize   = 1000
	DefaultSize = 10
)

func paginate(page, size *int, db *gorm.DB) *gorm.DB {
	if *page == 0 {
		*page = 1
	}

	switch {
	case *size >= MaxSize:
		*size = MaxSize
	case *size <= 0:
		*size = DefaultSize
	}

	offset := (*page - 1) * *size
	return db.Offset(offset).Limit(*size)
}

type ListPageInput struct {
	Size int `comment:"每页记录数"`
	Page int `comment:"页数"`
}
