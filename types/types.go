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
	GetTasksDetails() ([]TaskDetails, error)
	GetTaskById(id int) (*Task, error)
	GetTaskDetailsById(id int) (*TaskDetails, error)
	UpdateTaskById(id int, payload UpdateTaskPayload) error
	DeleteTaskById(id int) error
}

type UserStore interface {
	CreateUser(payload CreateUserPayload) error
	GetUsers() ([]User, error)
	GetUserById(id int) (*User, error)
	UpdateUserById(id int, payload UpdateUserPayload) error
	DeleteUserById(id int) error
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
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type TaskDetails struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	CreatedAt          string `json:"created_at"`
	Username           string `json:"username"`
	UserEmail          string `json:"user_email"`
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
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

type BaseUserPayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type CreateUserPayload struct {
	BaseUserPayload
}

type UpdateUserPayload struct {
	BaseUserPayload
}
