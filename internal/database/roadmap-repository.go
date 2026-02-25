package database

import (
	"database/sql"

	entity "github.com/J-V-S-C/MindBox/internal/entities"
	"github.com/google/uuid"
)

type RoadmapRepository struct {
	database *sql.DB
}

func NewRoadmapRepository(database *sql.DB) *RoadmapRepository {
	return &RoadmapRepository{database: database}
}

func (repository *RoadmapRepository) Create(name string, description string) (entity.Roadmap, error) {
	id := uuid.New().String()

	query := "INSERT INTO roadmaps (id, name, description) VALUES ($1, $2, $3)"
	_, err := repository.database.Exec(query, id, name, description)
	if err != nil {
		return entity.Roadmap{}, err
	}

	return entity.Roadmap{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (repository *RoadmapRepository) FindAll(limit int, offset int) ([]entity.Roadmap, error) {
	query := "SELECT id, name, description FROM roadmaps LIMIT $1 OFFSET $2"
	rows, err := repository.database.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roadmaps []entity.Roadmap
	for rows.Next() {
		var roadmap entity.Roadmap
		err := rows.Scan(
			&roadmap.ID,
			&roadmap.Name,
			&roadmap.Description,
		)
		if err != nil {
			return nil, err
		}
		roadmaps = append(roadmaps, roadmap)
	}

	return roadmaps, nil
}

func (repository *RoadmapRepository) FindByID(id string) (entity.Roadmap, error) {
	query := "SELECT id, name, description FROM roadmaps WHERE id = $1"
	row := repository.database.QueryRow(query, id)

	var roadmap entity.Roadmap
	err := row.Scan(&roadmap.ID, &roadmap.Name, &roadmap.Description)
	if err != nil {
		return entity.Roadmap{}, err
	}

	return roadmap, nil
}

func (repository *RoadmapRepository) Update(id string, name string, description string) (entity.Roadmap, error) {
	query := "UPDATE roadmaps SET name = $1, description = $2 WHERE id = $3"
	_, err := repository.database.Exec(query, name, description, id)
	if err != nil {
		return entity.Roadmap{}, err
	}

	return entity.Roadmap{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (repository *RoadmapRepository) Delete(id string) error {
	query := "DELETE FROM roadmaps WHERE id = $1"
	_, err := repository.database.Exec(query, id)
	return err
}