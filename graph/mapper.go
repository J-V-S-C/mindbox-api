package graph

import (
	"github.com/J-V-S-C/MindBox/graph/model"
	entity "github.com/J-V-S-C/MindBox/internal/entities"
)

func ToTaskModel(t entity.Task) *model.Task {
	return &model.Task{
		ID:          t.ID,
		Name:        t.Name,
		Description: &t.Description,
		Done:        t.Done,
		IsDaily:     t.IsDaily,
		Lifetime:    &t.Lifetime,
		IsExpired:   t.IsExpired,
	}
}

func ToCategoryModel(c entity.Category) *model.Category {
	return &model.Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: &c.Description,
		Lifetime:    c.Lifetime,
	}
}

func ToRoadmapModel(r entity.Roadmap) *model.Roadmap {
	return &model.Roadmap{
		ID:          r.ID,
		Name:        r.Name,
		Description: &r.Description,
	}
}