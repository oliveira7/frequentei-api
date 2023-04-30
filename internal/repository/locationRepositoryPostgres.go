package repository

import (
	"database/sql"
	"frequentei-api/internal/entity"
)

type LocationRepositoryPostgres struct {
	DB *sql.DB
}

func NewLocationRepositoryPostgres(db *sql.DB) *LocationRepositoryPostgres {
	return &LocationRepositoryPostgres{DB: db}
}

func (r *LocationRepositoryPostgres) Create(location *entity.Location) error {
	_, err := r.DB.Exec("INSERT INTO locations (id, name, description) VALUES (?, ?, ?)",
		location.ID, location.Name, location.Description)

	if err != nil {
		return err
	}
	return nil
}

func (r *LocationRepositoryPostgres) FindAll() ([]*entity.Location, error) {
	rows, err := r.DB.Query("SELECT * FROM locations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []*entity.Location

	for rows.Next() {
		var location entity.Location
		err := rows.Scan(&location.ID, &location.Name, &location.Description)
		if err != nil {
			return nil, err
		}
		locations = append(locations, &location)
	}

	return locations, nil
}
