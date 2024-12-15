package types

// interfaces types
type ProjectStore interface {
	CreateProject(payload CreateProjectPayload) error
	GetProjects() ([]Project, error)
	GetProjectById(id int) (*Project, error)
}

type TaskStore interface {
	CreateTask(payload CreateTaskPayload) error
}

// database models types
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type Task struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	ProjectID  int    `json:"project_id"`
	AssignedTo int    `json:"assigned_to"`
	CreatedAt  string `json:"created_at"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// http request payload types
type CreateProjectPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateTaskPayload struct {
	Name       string `json:"name" validate:"required"`
	Status     string `json:"status" validate:"required"`
	ProjectID  int    `json:"project_id" validate:"required"`
	AssignedTo int    `json:"assigned_to" validate:"required"`
}
