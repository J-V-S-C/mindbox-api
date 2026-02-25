package database

import (
	"database/sql"

	entity "github.com/J-V-S-C/MindBox/internal/entities"
	"github.com/google/uuid"
)

type CategoryRepository struct {
	database *sql.DB
}

func NewCategoryRepository(database *sql.DB) *CategoryRepository {
	return &CategoryRepository{database: database}
}

func (repository *CategoryRepository) Create(name string, description string, lifetime string, roadmapID string) (entity.Category, error) {
	id := uuid.New().String()

	query := "INSERT INTO categories (id, name, description, lifetime, roadmap_id) VALUES ($1, $2, $3, $4, $5)"
	_, err := repository.database.Exec(query, id, name, description, lifetime, roadmapID)
	if err != nil {
		return entity.Category{}, err
	}

	return entity.Category{
		ID:          id,
		Name:        name,
		Description: description,
		Lifetime:    lifetime,
		RoadmapID:   roadmapID,
	}, nil
}

func (repository *CategoryRepository) FindByID(id string) (entity.Category, error) {
	query := "SELECT id, name, description, lifetime, roadmap_id FROM categories WHERE id = $1"
	row := repository.database.QueryRow(query, id)

	var category entity.Category
	err := row.Scan(&category.ID, &category.Name, &category.Description, &category.Lifetime, &category.RoadmapID)
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (repository *CategoryRepository) FindByRoadmapID(roadmapID string) ([]entity.Category, error) {
	query := "SELECT id, name, description, lifetime, roadmap_id FROM categories WHERE roadmap_id = $1"
	rows, err := repository.database.Query(query, roadmapID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repository.scanCategories(rows)
}

func (repository *CategoryRepository) FindAll(limit int, offset int) ([]entity.Category, error) {
	query := "SELECT id, name, description, lifetime, roadmap_id FROM categories LIMIT $1 OFFSET $2"
	rows, err := repository.database.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repository.scanCategories(rows)
}

func (repository *CategoryRepository) Update(id string, name string, description string, lifetime string) (entity.Category, error) {
	query := `
		UPDATE categories 
		SET name = $1, description = $2, lifetime = $3 
		WHERE id = $4 
		RETURNING id, name, description, lifetime, roadmap_id
	`
	row := repository.database.QueryRow(query, name, description, lifetime, id)

	var category entity.Category
	err := row.Scan(&category.ID, &category.Name, &category.Description, &category.Lifetime, &category.RoadmapID)
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (repository *CategoryRepository) Delete(id string) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := repository.database.Exec(query, id)
	return err
}

func (repository *CategoryRepository) scanCategories(rows *sql.Rows) ([]entity.Category, error) {
	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.Lifetime,
			&category.RoadmapID,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}