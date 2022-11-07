package interfaces

import (
	paginate "iecare-api/src/app/pkg/paginate"
)

type BaseRepository[T interface{}] interface {
	List(meta paginate.Meta) (*paginate.Pagination, error)
	Get(id string) (*T, error)
	Store(model *T) (*T, error)
	Edit(model *T) (*T, error)
	Delete(model *T) error
	FindBy(field string, value string) (*T, error)
	FindByMany(field []string, value string) (*T, error)
}
