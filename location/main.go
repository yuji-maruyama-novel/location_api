package location

import "gorm.io/gorm"

var (
	repository *LocationRepository
	service    *LocationService
)

func Location(db *gorm.DB) *LocationService {
	repository = newLocationRepository(db)
	service = newLocationService(repository)
	return service
}
