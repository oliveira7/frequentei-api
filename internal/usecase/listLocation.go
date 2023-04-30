package usecase

import "frequentei-api/internal/entity"

type ListLocationOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ListLocationUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewListLocationUseCase(locationRepository entity.LocationRepository) *ListLocationUseCase {
	return &ListLocationUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *ListLocationUseCase) Execute() ([]*ListLocationOutputDto, error) {
	locations, err := u.LocationRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var locationsOutput []*ListLocationOutputDto
	for _, location := range locations {
		locationsOutput = append(locationsOutput, &ListLocationOutputDto{
			ID:          location.ID,
			Name:        location.Name,
			Description: location.Description,
		})
	}
	return locationsOutput, nil
}
