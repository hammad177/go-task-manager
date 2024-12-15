package tasks

import (
	"database/sql"

	"github.com/hammad177/task_management/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateTask(payload types.CreateTaskPayload) error {
	_, err := s.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)", payload.Name, payload.Status, payload.ProjectID, payload.AssignedTo)

	if err != nil {
		return err
	}

	return nil
}
