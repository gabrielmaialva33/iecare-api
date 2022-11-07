package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	"iecare-api/src/app/repositories"
)

type Repositories struct {
	User     interfaces.UserInterface
	Role     interfaces.RoleInterface
	Provider interfaces.ProviderInterface
	Category interfaces.CategoryInterface
	db       *gorm.DB
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
		User:     repositories.NewUserRepository(database),
		Role:     repositories.NewRoleRepository(database),
		Provider: repositories.NewProvidersRepository(database),
		Category: repositories.NewCategoriesRepository(database),
		db:       database,
	}
}

func (r *Repositories) Migrate() {
	r.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := r.db.AutoMigrate(&models.User{}, &models.Role{}, &models.UserRole{}, &models.Provider{}, &models.Category{}); err != nil {
		panic(" -> Could not migrate the database")
	}

	err := r.db.SetupJoinTable(&models.User{}, "Roles", &models.UserRole{})
	if err != nil {
		panic(" -> Could not setup join table")
	}
}

func (r *Repositories) Drop() {
	r.db.Exec("DROP EXTENSION IF EXISTS \"uuid-ossp\" CASCADE;")
	err := r.db.Migrator().DropTable(&models.User{}, &models.Role{}, &models.UserRole{}, &models.Provider{}, &models.Category{})
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
			Name:        "provider",
			Slug:        "Provider",
			Description: "A provider user has all permissions except root and admin",
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
			Email:     "root@iecare.com",
			UserName:  "root",
			Password:  "123456",
			Role:      models.RoleRoot,
		},
		{
			FirstName: "Admin",
			LastName:  "System",
			Email:     "admin@iecare.com",
			UserName:  "admin",
			Password:  "123456",
			Role:      models.RoleAdmin,
		},
		{
			FirstName: "IECare",
			LastName:  "System",
			Email:     "iecare@iecare.com",
			UserName:  "iecare",
			Password:  "123456",
			Role:      models.RoleProvider,
		},
		{
			FirstName: "Gabriel",
			LastName:  "Maia",
			Email:     "maia@iecare.com",
			UserName:  "maia",
			Password:  "123456",
			Role:      models.RoleUser,
		},
		{
			FirstName: "Guest",
			LastName:  "System",
			Email:     "guest@iecare.com",
			UserName:  "guest",
			Password:  "123456",
			Role:      models.RoleGuest,
		},
	}
	categories := []models.Category{
		{
			Name:        "Encanamento",
			Description: "Tudo relacionado a encanamento",
			Icon:        "shower",
		},
		{
			Name:        "Elétrica",
			Description: "Tudo relacionado a elétrica",
			Icon:        "bolt",
		},
		{
			Name:        "Pintura",
			Description: "Tudo relacionado a pintura",
			Icon:        "paint-brush",
		},
		{
			Name:        "Jardinagem",
			Description: "Tudo relacionado a jardinagem",
			Icon:        "tree",
		},
		{
			Name:        "Limpeza",
			Description: "Tudo relacionado a limpeza",
			Icon:        "broom",
		},
		{
			Name:        "Marcenaria",
			Description: "Tudo relacionado a marcenaria",
			Icon:        "tools",
		},
		{
			Name:        "Mecânica",
			Description: "Tudo relacionado a mecânica",
			Icon:        "car",
		},
		{
			Name:        "Informática",
			Description: "Tudo relacionado a informática",
			Icon:        "laptop",
		},
		{
			Name:        "Montagem",
			Description: "Tudo relacionado a montagem",
			Icon:        "wrench",
		},
		{
			Name:        "Reformas",
			Description: "Tudo relacionado a reformas",
			Icon:        "hammer",
		},
		{
			Name:        "Outros",
			Description: "Outros serviços",
			Icon:        "question",
		},
	}

	r.db.Create(&roles)
	r.db.Create(&users)
	r.db.Create(&categories)
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
