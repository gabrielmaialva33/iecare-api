package interfaces

import (
	pagination2 "iecare-api/src/app/pkg/pagination"
)

type BaseRepository[T interface{}] interface {
	List(meta pagination2.Meta) (*pagination2.Pagination, error)
	Get(id string) (*T, error)
	Store(model *T) (*T, error)
	Edit(model *T) (*T, error)
	Delete(model *T) error
	FindBy(field string, value string) (*T, error)
	FindManyBy(field []string, value string) (*T, error)
}
