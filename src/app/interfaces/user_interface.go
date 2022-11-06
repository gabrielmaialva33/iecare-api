package interfaces

import (
	"iecare-api/src/app/models"
)

type UserInterface interface {
	BaseRepository[models.User]
}
