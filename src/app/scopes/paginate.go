package scopes

import (
	"gorm.io/gorm"
	"iecare-api/src/app/pkg/pagination"
	"math"
)

func Paginate(model interface{}, fields []string, meta *pagination.Meta, db *gorm.DB) func(db *gorm.DB) *gorm.DB {

	var count int64
	for _, field := range fields {
		db = db.Or(field+" ilike ?", "%"+meta.GetSearch()+"%")
	}

	db.Model(model).Count(&count)

	meta.Total = count
	meta.TotalPages = int(math.Ceil(float64(count) / float64(meta.GetPerPage())))
	meta.CurrentPage = meta.GetCurrentPage()
	meta.PerPage = meta.GetPerPage()
	meta.Sort = meta.GetSort(fields)
	meta.Order = meta.GetOrder()
	meta.Search = meta.GetSearch()

	return func(db *gorm.DB) *gorm.DB {
		for _, field := range fields {
			db = db.Or(field+" ilike ?", "%"+meta.GetSearch()+"%")
		}
		return db.Offset(meta.GetOffset()).Limit(meta.GetPerPage()).Order(meta.GetSort(fields))
	}
}
