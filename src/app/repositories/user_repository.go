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

type UserRepo struct {
	db *gorm.DB
}

// NewUserRepository returns a new instance of UserRepo
func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

// UserRepo implements interfaces.UserInterface
var _ interfaces.UserInterface = &UserRepo{}

func (u *UserRepo) List(meta pagination2.Meta) (*pagination2.Pagination, error) {
	var users models.Users
	var fields = []string{"first_name", "last_name", "email", "user_name"}
	var paginate pagination2.Pagination

	if err := u.db.Preload("Roles").Scopes(scopes.Paginate(users, fields, &meta, u.db)).Find(&users).Error; err != nil {
		return nil, err
	}

	paginate.SetMeta(meta)
	paginate.SetData(users.PublicUsers())

	return &paginate, nil
}

func (u *UserRepo) Get(id string) (*models.User, error) {
	var user models.User
	if err := u.db.Preload("Roles").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) Store(user *models.User) (*models.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}
	u.db.Preload("Roles").Where("id = ?", user.Id).First(&user)
	return user, nil
}

func (u *UserRepo) Edit(user *models.User) (*models.User, error) {
	if err := u.db.Clauses(clause.Returning{}).Preload("Roles").Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) Delete(user *models.User) error {
	u.db.Where("id = ?", user.Id).Updates(&user)
	if err := u.db.Where("id = ?", user.Id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) FindBy(field string, value string) (*models.User, error) {
	var user models.User
	if err := u.db.Preload("Roles").Where(field+" = ?", value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) FindManyBy(field []string, value string) (*models.User, error) {
	var user models.User
	for _, f := range field {
		u.db.Preload("Roles").Where(f+" = ?", value).First(&user)
		if user.Id != "" {
			return &user, nil
		}
	}
	return nil, errors.New("record not found")
}
