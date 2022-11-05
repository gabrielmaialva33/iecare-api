package interfaces

import (
	"iecare-api/src/app/modules/accounts/models"
	"iecare-api/src/app/shared/interfaces"
)

type UserInterface interface {
	interfaces.BaseRepository[models.User]
}
