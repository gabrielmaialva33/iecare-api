package interfaces

import "iecare-api/src/app/models"

type CategoryInterface interface {
	BaseRepository[models.Category]
}
