package interfaces

import (
	"iecare-api/src/app/models"
)

type ProviderInterface interface {
	BaseRepository[models.Provider]
}
