package models

import (
	"gorm.io/gorm"
	"iecare-api/src/app/pkg"
	"time"
)

type User struct {
	gorm.Model

	Id        string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id"`
	Avatar    string         `gorm:"column:avatar;type:text;default:null;" json:"avatar" validate:"omitempty,url"`
	FirstName string         `gorm:"column:first_name;size:80;not null;" json:"first_name" validate:"required,min=2,max=80,alpha,omitempty"`
	LastName  string         `gorm:"column:last_name;size:80;not null;" json:"last_name" validate:"required,min=2,max=80,alpha,omitempty"`
	FullName  string         `gorm:"->;type:varchar(160) generated always as (first_name || ' ' || last_name) stored;default:(-);" json:"full_name"`
	Email     string         `gorm:"column:email;size:255;not null;unique;unique;index;" json:"email" validate:"required,email,max=247,unique,omitempty"`
	UserName  string         `gorm:"column:user_name;size:58;not null;unique;index;" json:"user_name" validate:"required,min=4,max=50,unique,omitempty"`
	Password  string         `gorm:"column:password;size:255;not null;" json:"password" form:"password" validate:"required,min=6,max=50,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index,default:null;" json:"-"`

	// Relationships
	Roles Roles `gorm:"many2many:user_roles;" json:"roles"`

	// Virtual Fields
	Role            string `gorm:"-" json:"-"`
	ConfirmPassword string `gorm:"-" json:"confirm_password" form:"confirm_password" validate:"required,min=6,max=50,eqfield=Password,omitempty"`
}

type Users []User

type UserPublic struct {
	Id        string `json:"id"`
	Avatar    string `json:"avatar"`
	FullName  string `json:"full_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`

	Roles []interface{} `json:"roles"`
}

// PublicUser returns a public representation of a user.
func (u *User) PublicUser() interface{} {
	return &UserPublic{
		Id:        u.Id,
		Avatar:    u.Avatar,
		FullName:  u.FullName,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		UserName:  u.UserName,
		Roles:     u.Roles.PublicRoles(),
	}
}

// PublicUsers returns a public representation of a user.
func (u Users) PublicUsers() []interface{} {
	users := make([]interface{}, len(u))
	for i, user := range u {
		users[i] = user.PublicUser()
	}
	if len(users) == 0 {
		return []interface{}{}
	}
	return users
}

// BeforeSave hook executed before saving a User to the database.
func (u *User) BeforeSave(*gorm.DB) error {
	hash, err := pkg.CreateHash(u.Password, pkg.DefaultParams)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) AfterCreate(db *gorm.DB) error {
	var role Role
	if u.Role == "" {
		db.Where("name = ?", RoleGuest).First(&role)
	} else if u.Role == RoleAdmin {
		db.Where("name = ?", RoleAdmin).First(&role)
	} else if u.Role == RoleUser {

		db.Where("name = ?", RoleUser).First(&role)
	} else if u.Role == RoleProvider {
		db.Where("name = ?", RoleProvider).First(&role)
	} else if u.Role == RoleRoot {
		db.Where("name = ?", RoleRoot).First(&role)
	} else {
		db.Where("name = ?", RoleGuest).First(&role)
	}

	err := db.Model(&u).Association("Roles").Append(&role)
	if err != nil {
		return err
	}
	return nil
}

type Login struct {
	Uid      string `json:"uid" validate:"required"`
	Password string `json:"password" validate:"required"`
}
