package location

import (
	"gqlgen/graph/model"
)

type LocationService struct {
	locationRepository LocationInterface
}

func newLocationService(locationRepository LocationInterface) *LocationService {
	return &LocationService{locationRepository: locationRepository}
}

func (ls *LocationService) CreateLocation(input model.NewLocation) (*model.Location, error) {
	return ls.locationRepository.createLocation(input)
}

func (ls *LocationService) GetLocations() ([]*model.Location, error) {
	return ls.locationRepository.getLocations()
}
