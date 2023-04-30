package entity

import (
	"github.com/google/uuid"
)

type LocationRepository interface {
	FindAll() ([]*Location, error)
	Create(location *Location) error
}

type Location struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewLocation(name string, description string) *Location {
	return &Location{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
	}
}
