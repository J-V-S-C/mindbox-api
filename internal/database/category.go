package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db *sql.DB
	ID string
	Name string
	Description string
	RoadmapID string
}

func NewCategory(db *sql.DB) *Category{
	return &Category{db: db}
}

func (category *Category) Create(name string, description string, roadmapId string) (Category, error){
	id := uuid.New().String()
	_, err := category.db.Exec("INSERT INTO categories (id, name, description, roadmap_id) VALUES ($1, $2, $3, $4)", id, name, description, roadmapId)
	if err != nil {
		return Category{}, err 
	}

	return Category{ID: id, Name: name, Description: description, RoadmapID: roadmapId}, nil
}