package interfaces

import (
	"iecare-api/src/app/models"
)

type RoleInterface interface {
	BaseRepository[models.Role]
}
