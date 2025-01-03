package tasks

import (
	"database/sql"
	"fmt"

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

func (s *Store) GetTasks() ([]types.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	tasks := make([]types.Task, 0)

	for rows.Next() {
		t, err := scanRowsIntoTasks(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *t)
	}

	return tasks, nil
}

func (s *Store) GetTasksDetails() ([]types.TaskDetails, error) {
	query := `
		SELECT 
			t.id, t.name, t.status, t.created_at, 
			u.username, u.email as user_email,
			p.name as project_name, p.description as project_description
		FROM 
			tasks t
		LEFT JOIN 
			users u ON t.assigned_to = u.id
		LEFT JOIN 
			projects p ON t.project_id = p.id
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	tasks := make([]types.TaskDetails, 0)

	for rows.Next() {
		t, err := scanRowsIntoTasksDetails(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *t)
	}

	return tasks, nil
}

func (s *Store) GetTaskById(id int) (*types.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	p := new(types.Task)
	for rows.Next() {
		p, err = scanRowsIntoTasks(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, fmt.Errorf("task not found")
	}

	return p, nil
}

func (s *Store) GetTaskDetailsById(id int) (*types.TaskDetails, error) {
	query := `
		SELECT 
			t.id, t.name, t.status, t.created_at, 
			u.username, u.email as user_email,
			p.name as project_name, p.description as project_description
		FROM 
			tasks t
		LEFT JOIN 
			users u ON t.assigned_to = u.id
		LEFT JOIN 
			projects p ON t.project_id = p.id
		WHERE 
			t.id = ?
	`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	t := new(types.TaskDetails)
	for rows.Next() {
		t, err = scanRowsIntoTasksDetails(rows)
		if err != nil {
			return nil, err
		}
	}

	if t.ID == 0 {
		return nil, fmt.Errorf("task not found")
	}

	return t, nil
}

func (s *Store) UpdateTaskById(id int, payload types.UpdateTaskPayload) error {
	_, err1 := s.GetTaskById(id)

	if err1 != nil {
		return err1
	}

	_, err2 := s.db.Exec("UPDATE tasks SET name = ?, status = ?, project_id = ?, assigned_to = ? WHERE id = ?", payload.Name, payload.Status, payload.ProjectID, payload.AssignedTo, id)

	if err2 != nil {
		return err2
	}

	return nil
}

func (s *Store) DeleteTaskById(id int) error {
	_, err1 := s.GetTaskById(id)

	if err1 != nil {
		return err1
	}

	_, err2 := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err2 != nil {
		return err2
	}

	return nil
}

func scanRowsIntoTasks(rows *sql.Rows) (*types.Task, error) {
	t := new(types.Task)
	err := rows.Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedTo, &t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func scanRowsIntoTasksDetails(rows *sql.Rows) (*types.TaskDetails, error) {
	t := new(types.TaskDetails)
	err := rows.Scan(&t.ID, &t.Name, &t.Status, &t.CreatedAt, &t.Username, &t.UserEmail, &t.ProjectName, &t.ProjectDescription)
	if err != nil {
		return nil, err
	}
	return t, nil
}
