package usecase

import "frequentei-api/internal/entity"

type CreateLocationInputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateLocationOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateLocationUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewCreateLocationUseCase(locationRepository entity.LocationRepository) *CreateLocationUseCase {
	return &CreateLocationUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *CreateLocationUseCase) Execute(input CreateLocationInputDto) (*CreateLocationOutputDto, error) {
	location := entity.NewLocation(input.Name, input.Description)

	err := u.LocationRepository.Create(location)
	if err != nil {
		return nil, err
	}

	return &CreateLocationOutputDto{
		ID:          location.ID,
		Name:        location.Name,
		Description: location.Description,
	}, nil
}
