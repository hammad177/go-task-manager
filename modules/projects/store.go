package projects

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

func (s *Store) CreateProject(payload types.CreateProjectPayload) error {
	_, err := s.db.Exec("INSERT INTO projects (name, description) VALUES (?, ?)", payload.Name, payload.Description)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetProjects() ([]types.Project, error) {
	rows, err := s.db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}

	projects := make([]types.Project, 0)

	for rows.Next() {
		p, err := scanRowsIntoProject(rows)
		if err != nil {
			return nil, err
		}
		projects = append(projects, *p)
	}

	return projects, nil
}

func (s *Store) GetProjectById(id int) (*types.Project, error) {
	rows, err := s.db.Query("SELECT * FROM projects WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	p := new(types.Project)
	for rows.Next() {
		p, err = scanRowsIntoProject(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, fmt.Errorf("project not found")
	}

	return p, nil
}

func scanRowsIntoProject(rows *sql.Rows) (*types.Project, error) {
	p := new(types.Project)
	err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}
