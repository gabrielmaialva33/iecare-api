package repositories

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	pagination2 "iecare-api/src/app/pkg/pagination"
	"iecare-api/src/app/scopes"
)

type RoleRepo struct {
	db *gorm.DB
}

// NewRoleRepository returns a new instance of RoleRepo
func NewRoleRepository(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db}
}

// RoleRepo implements interfaces.RoleInterface
var _ interfaces.RoleInterface = &RoleRepo{}

func (r *RoleRepo) List(meta pagination2.Meta) (*pagination2.Pagination, error) {
	var roles models.Roles
	var fields = []string{"name", "slug", "description"}
	var paginate pagination2.Pagination

	if err := r.db.Scopes(scopes.Paginate(roles, fields, &meta, r.db)).Find(&roles).Error; err != nil {
		return nil, err
	}

	paginate.SetMeta(meta)
	paginate.SetData(roles.PublicRoles())

	return &paginate, nil
}

func (r *RoleRepo) Get(id string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepo) Store(role *models.Role) (*models.Role, error) {
	if err := r.db.Create(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepo) Edit(role *models.Role) (*models.Role, error) {
	if err := r.db.Clauses(clause.Returning{}).Where("id = ?", role.Id).Updates(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepo) Delete(role *models.Role) error {
	r.db.Where("id = ?", role.Id).Updates(&role)
	if err := r.db.Where("id = ?", role.Id).Delete(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoleRepo) FindBy(field string, value string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where(field+" = ?", value).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepo) FindManyBy(field []string, value string) (*models.Role, error) {
	var role models.Role
	for _, f := range field {
		r.db.Where(f+" = ?", value).First(&role)
		if role.Id != "" {
			return &role, nil
		}
	}
	return nil, errors.New("user not found")
}
