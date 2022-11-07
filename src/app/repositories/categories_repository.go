package repositories

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	"iecare-api/src/app/pkg/paginate"
	"iecare-api/src/app/scopes"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db}
}

var _ interfaces.BaseRepository[models.Category] = &CategoryRepo{}

func (c CategoryRepo) List(meta paginate.Meta) (*paginate.Pagination, error) {
	var categories models.Categories
	var fields = []string{"name"}
	var pagination paginate.Pagination

	if err := c.db.Scopes(scopes.Paginate(categories, fields, &meta, c.db)).Find(&categories).Error; err != nil {
		return nil, err
	}

	pagination.SetMeta(meta)
	pagination.SetData(categories.PublicCategories())

	return &pagination, nil

}

func (c CategoryRepo) Get(id string) (*models.Category, error) {
	var category models.Category
	if err := c.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepo) Store(category *models.Category) (*models.Category, error) {
	if err := c.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c CategoryRepo) Edit(category *models.Category) (*models.Category, error) {
	if err := c.db.Clauses(clause.Returning{}).Where("id = ?", category.Id).Updates(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c CategoryRepo) Delete(category *models.Category) error {
	c.db.Where("id = ?", category.Id).Updates(&category)
	if err := c.db.Where("id = ?", category.Id).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func (c CategoryRepo) FindBy(field string, value string) (*models.Category, error) {
	var category models.Category
	if err := c.db.Where(field+" = ?", value).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepo) FindByMany(field []string, value string) (*models.Category, error) {
	var category models.Category
	for _, f := range field {
		c.db.Where(f+" = ?", value).First(&category)
		if category.Id != "" {
			return &category, nil
		}
	}
	return nil, errors.New("record not found")
}
