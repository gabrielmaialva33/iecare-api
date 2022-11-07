package services

import (
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	"iecare-api/src/app/pkg/paginate"
)

type CategoryServices struct {
	cr interfaces.CategoryInterface
}

type ICategoryServices interface {
	interfaces.CategoryInterface
}

var _ ICategoryServices = &CategoryServices{}

func (c CategoryServices) List(meta paginate.Meta) (*paginate.Pagination, error) {
	return c.cr.List(meta)
}

func (c CategoryServices) Get(id string) (*models.Category, error) {
	return c.cr.Get(id)
}

func (c CategoryServices) Store(category *models.Category) (*models.Category, error) {
	return c.cr.Store(category)
}

func (c CategoryServices) Edit(category *models.Category) (*models.Category, error) {
	return c.cr.Edit(category)
}

func (c CategoryServices) Delete(category *models.Category) error {
	return c.cr.Delete(category)
}

func (c CategoryServices) FindBy(field string, value string) (*models.Category, error) {
	return c.cr.FindBy(field, value)
}

func (c CategoryServices) FindByMany(field []string, value string) (*models.Category, error) {
	return c.cr.FindByMany(field, value)
}
