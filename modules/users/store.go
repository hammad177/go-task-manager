package users

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

func (s *Store) CreateUser(payload types.CreateUserPayload) error {
	_, err := s.db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", payload.Username, payload.Email)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUsers() ([]types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users := make([]types.User, 0)

	for rows.Next() {
		u, err := scanRowsIntoUsers(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}

	return users, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUsers(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) UpdateUserById(id int, payload types.UpdateUserPayload) error {
	_, err1 := s.GetUserById(id)

	if err1 != nil {
		return err1
	}

	_, err2 := s.db.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", payload.Username, payload.Email, id)

	if err2 != nil {
		return err2
	}

	return nil
}

func (s *Store) DeleteUserById(id int) error {
	_, err1 := s.GetUserById(id)

	if err1 != nil {
		return err1
	}

	_, err2 := s.db.Exec("DELETE FROM users WHERE id = ?", id)

	if err2 != nil {
		return err2
	}

	return nil
}

func scanRowsIntoUsers(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
	err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
