package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Roadmap struct {
	db *sql.DB
	ID string
	Name string
	Description string
}

func NewRoadmap(db *sql.DB) *Roadmap{
	return &Roadmap{db: db}
}

func (roadmap *Roadmap) Create(name string, description string) (Roadmap, error){
	id := uuid.New().String()
	_, err := roadmap.db.Exec("INSERT INTO roadmaps (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return Roadmap{}, err 
	}

	return Roadmap{ID: id, Name: name, Description: description}, nil
}

func (roadmap *Roadmap) FindAll() ([]Roadmap, error){
	rows, err := roadmap.db.Query("SELECT id, name, description FROM roadmaps")
	if err != nil {
		return nil, err 
	}
	defer rows.Close()
	roadmaps := []Roadmap{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		roadmaps = append(roadmaps, Roadmap{ID: id, Name: name, Description: description})
	}
	return roadmaps, nil
}