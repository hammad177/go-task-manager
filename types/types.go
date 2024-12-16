package types

// interfaces types
type ProjectStore interface {
	CreateProject(payload CreateProjectPayload) error
	GetProjects() ([]Project, error)
	GetProjectById(id int) (*Project, error)
	UpdateProjectById(id int, payload UpdateProjectPayload) error
	DeleteProjectById(id int) error
}

type TaskStore interface {
	CreateTask(payload CreateTaskPayload) error
	GetTasks() ([]Task, error)
	GetTaskById(id int) (*Task, error)
	UpdateTaskById(id int, payload UpdateTaskPayload) error
	DeleteTaskById(id int) error
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
type BaseProjectPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateProjectPayload struct {
	BaseProjectPayload
}

type UpdateProjectPayload struct {
	BaseProjectPayload
}
type BaseTaskPayload struct {
	Name       string `json:"name" validate:"required"`
	Status     string `json:"status" validate:"required"`
	ProjectID  int    `json:"project_id" validate:"required"`
	AssignedTo int    `json:"assigned_to" validate:"required"`
}

type CreateTaskPayload struct {
	BaseTaskPayload
}

type UpdateTaskPayload struct {
	BaseTaskPayload
}
