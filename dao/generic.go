package dao

import (
	"errors"
	"math"

	"gorm.io/gorm"
)

func GInsert[T any](db *gorm.DB, data T) error {
	return db.Create(data).Error
}

func GDelete[T any](db *gorm.DB, query string, args ...any) error {
	var t T
	return db.Model(t).Where(query, args...).Delete(&t).Error
}

func GPaginateOrder[T any](db *gorm.DB, params *ListPageInput, order, query string, args ...any) ([]T, int64, int64, error) {
	var t T
	var results []T
	var count int64

	page := params.Page
	size := params.Size
	q := db.Where(query, args...)

	if order != "" {
		q = q.Order(order)
	}

	err := paginate(&page, &size, q).Find(&results).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, 0, err
	}

	err = db.Model(&t).Where(query, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, 0, err
	}

	pages := int64(math.Ceil(float64(count) / float64(size)))

	return results, count, pages, nil
}

func GWhereFirstSelect[T any](db *gorm.DB, columns string, query string, args ...any) (*T, error) {
	var result T
	q := db.Model(&result)

	if columns != "" && columns != "*" {
		q = q.Select(columns)
	}
	q = q.Where(query, args...)
	err := q.First(&result).Error
	return &result, err
}

func GWhereAllSelectOrder[T any](db *gorm.DB, columns string, order string, query string, args ...any) ([]T, error) {
	return gWhereAllSelectOrderLimit[T](db, columns, order, 0, query, args...)
}

func gWhereAllSelectOrderLimit[T any](db *gorm.DB, columns string, order string, limit int, query string, args ...any) ([]T, error) {
	var results []T
	var t T
	q := db.Model(t)

	if order != "" {
		q = q.Order(order)
	}

	if columns != "" && columns != "*" {
		q = q.Select(columns)
	}
	q = q.Where(query, args...)
	if limit > 0 {
		q = q.Limit(limit)
	}
	err := q.Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
