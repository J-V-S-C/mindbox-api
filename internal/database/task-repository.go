package database

import (
	"database/sql"

	entity "github.com/J-V-S-C/MindBox/internal/entities"
	"github.com/google/uuid"
)

type TaskRepository struct {
	database *sql.DB
}

func NewTaskRepository(database *sql.DB) *TaskRepository {
	return &TaskRepository{database: database}
}

func (repository *TaskRepository) Create(name string, description string, isDaily bool, lifetime string, categoryID string) (entity.Task, error) {
	id := uuid.New().String()
	query := `
		INSERT INTO tasks (id, name, description, done, is_daily, lifetime, category_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := repository.database.Exec(query, id, name, description, false, isDaily, lifetime, categoryID)
	if err != nil {
		return entity.Task{}, err
	}

	return entity.Task{
		ID:          id,
		Name:        name,
		Description: description,
		Done:        false,
		IsDaily:     isDaily,
		Lifetime:    lifetime,
		CategoryID:  categoryID,
	}, nil
}

func (repository *TaskRepository) FindByID(id string) (entity.Task, error) {
	query := "SELECT id, name, description, done, is_daily, lifetime, category_id FROM tasks WHERE id = $1"
	row := repository.database.QueryRow(query, id)

	var task entity.Task
	err := row.Scan(
		&task.ID,
		&task.Name,
		&task.Description,
		&task.Done,
		&task.IsDaily,
		&task.Lifetime,
		&task.CategoryID,
	)
	if err != nil {
		return entity.Task{}, err
	}

	task.IsExpired = task.CheckExpired()
	return task, nil
}

func (repository *TaskRepository) FindAll(limit int, offset int) ([]entity.Task, error) {
	query := "SELECT id, name, description, done, is_daily, lifetime, category_id FROM tasks LIMIT $1 OFFSET $2"
	rows, err := repository.database.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repository.scanTasks(rows)
}

func (repository *TaskRepository) FindDailyTasks(limit int, offset int) ([]entity.Task, error) {
	query := "SELECT id, name, description, done, is_daily, lifetime, category_id FROM tasks WHERE is_daily = true LIMIT $1 OFFSET $2"
	rows, err := repository.database.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repository.scanTasks(rows)
}

func (repository *TaskRepository) FindByCategoryID(categoryID string, limit int, offset int) ([]entity.Task, error) {
	query := "SELECT id, name, description, done, is_daily, lifetime, category_id FROM tasks WHERE category_id = $1 LIMIT $2 OFFSET $3"
	rows, err := repository.database.Query(query, categoryID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repository.scanTasks(rows)
}

func (repository *TaskRepository) FindPendingByCategoryID(categoryID string, limit int, offset int) ([]entity.Task, error) {
	var rows *sql.Rows
	var err error

	query := "SELECT id, name, description, done, is_daily, lifetime, category_id FROM tasks WHERE done = false AND category_id = $1 LIMIT $2 OFFSET $3"
	rows, err = repository.database.Query(query, categoryID, limit, offset)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repository.scanTasks(rows)
}

func (repository *TaskRepository) FindExpiredTasks(limit int, offset int) ([]entity.Task, error) {
    query := "SELECT id, name, description, done, is_daily, lifetime, category_id FROM tasks WHERE lifetime != '' AND CAST(lifetime AS TIMESTAMP) < NOW() LIMIT $1 OFFSET $2"
    
    rows, err := repository.database.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return repository.scanTasks(rows)
}

func (repository *TaskRepository) ToggleDone(id string) (entity.Task, error) {
	query := "UPDATE tasks SET done = NOT done WHERE id = $1 RETURNING id, name, description, done, is_daily, lifetime, category_id"
	row := repository.database.QueryRow(query, id)

	var task entity.Task
	err := row.Scan(&task.ID, &task.Name, &task.Description, &task.Done, &task.IsDaily, &task.Lifetime, &task.CategoryID)
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (repository *TaskRepository) Update(id string, name string, description string, isDaily bool, lifetime string) (entity.Task, error) {
	query := `
		UPDATE tasks 
		SET name = $1, description = $2, is_daily = $3, lifetime = $4 
		WHERE id = $5 
		RETURNING id, name, description, done, is_daily, lifetime, category_id
	`
	row := repository.database.QueryRow(query, name, description, isDaily, lifetime, id)

	var task entity.Task
	err := row.Scan(
		&task.ID,
		&task.Name,
		&task.Description,
		&task.Done,
		&task.IsDaily,
		&task.Lifetime,
		&task.CategoryID,
	)
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (repository *TaskRepository) Delete(id string) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := repository.database.Exec(query, id)
	return err
}

func (repository *TaskRepository) scanTasks(rows *sql.Rows) ([]entity.Task, error) {
	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Description,
			&task.Done,
			&task.IsDaily,
			&task.Lifetime,
			&task.CategoryID,
		)
		if err != nil {
			return nil, err
		}

		task.IsExpired = task.CheckExpired()
		tasks = append(tasks, task)
	}
	return tasks, nil
}