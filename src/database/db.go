package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"iecare-api/src/app/modules/accounts/interfaces"
	"iecare-api/src/app/modules/accounts/models"
	"iecare-api/src/app/modules/accounts/repositories"
)

type Repositories struct {
	User interfaces.UserInterface
	Role interfaces.RoleInterface
	db   *gorm.DB
}

var DB *gorm.DB

func Connect(dsn string) *Repositories {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(" -> Could not connect to the database")
	}

	DB = database

	return &Repositories{
		User: repositories.NewUserRepository(database),
		Role: repositories.NewRoleRepository(database),
		db:   database,
	}
}

func (r *Repositories) Migrate() {
	r.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := r.db.AutoMigrate(&models.User{}, &models.Role{}, &models.UserRole{}); err != nil {
		panic(" -> Could not migrate the database")
	}

	err := r.db.SetupJoinTable(&models.User{}, "Roles", &models.UserRole{})
	if err != nil {
		panic(" -> Could not setup join table")
	}
}

func (r *Repositories) Drop() {
	r.db.Exec("DROP EXTENSION IF EXISTS \"uuid-ossp\" CASCADE;")
	err := r.db.Migrator().DropTable(&models.User{}, &models.Role{}, &models.UserRole{})
	if err != nil {
		panic(" -> Could not drop the database")
	}
}

func (r *Repositories) Seed() {
	roles := []models.Role{
		{
			Name:        "root",
			Slug:        "Root",
			Description: "A root user has all permissions",
		},
		{
			Name:        "admin",
			Slug:        "Admin",
			Description: "An admin user has all permissions except root",
		},
		{
			Name:        "user",
			Slug:        "User",
			Description: "A user has limited permissions",
		},
		{
			Name:        "guest",
			Slug:        "Guest",
			Description: "A guest user has no permissions",
		},
	}

	users := []models.User{
		{
			FirstName: "Root",
			LastName:  "System",
			Email:     "root@go.com",
			UserName:  "root",
			Password:  "123456",
			Role:      models.RoleRoot,
		},
		{
			FirstName: "Admin",
			LastName:  "System",
			Email:     "admin@go.com",
			UserName:  "admin",
			Password:  "123456",
			Role:      models.RoleAdmin,
		},
		{
			FirstName: "Gabriel",
			LastName:  "Maia",
			Email:     "maia@go.com",
			UserName:  "maia",
			Password:  "123456",
			Role:      models.RoleUser,
		},
		{
			FirstName: "Guest",
			LastName:  "System",
			Email:     "guest@go.com",
			UserName:  "guest",
			Password:  "123456",
			Role:      models.RoleGuest,
		},
	}

	r.db.Create(&roles)
	r.db.Create(&users)
}

func (r *Repositories) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (r *Repositories) Stats() map[string]interface{} {
	sqlDB, _ := r.db.DB()
	return map[string]interface{}{
		"open_connections":     sqlDB.Stats().OpenConnections,
		"max_open_connections": sqlDB.Stats().MaxOpenConnections,
		"in_use":               sqlDB.Stats().InUse,
		"idle":                 sqlDB.Stats().Idle,
		"max_idle_closed":      sqlDB.Stats().MaxIdleClosed,
		"max_idle_time_closed": sqlDB.Stats().MaxIdleTimeClosed,
		"max_lifetime_closed":  sqlDB.Stats().MaxLifetimeClosed,
		"wait_count":           sqlDB.Stats().WaitCount,
		"wait_duration":        sqlDB.Stats().WaitDuration,
		"ping":                 sqlDB.Ping(),
	}
}
