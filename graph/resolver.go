package graph

import "github.com/J-V-S-C/MindBox/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct{
	RoadmapDB *database.RoadmapRepository
	CategoryDB *database.CategoryRepository
	TaskDB *database.TaskRepository
}
