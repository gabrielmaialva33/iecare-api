package interfaces

import (
	"iecare-api/src/app/modules/accounts/models"
	"iecare-api/src/app/shared/interfaces"
)

type RoleInterface interface {
	interfaces.BaseRepository[models.Role]
}
