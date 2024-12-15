package tasks

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammad177/task_management/types"
	"github.com/hammad177/task_management/utils"
)

type Handler struct {
	store types.TaskStore
}

func NewHandler(store types.TaskStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", h.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks", h.GetTasks).Methods(http.MethodGet)
	// r.HandleFunc("/tasks/{id}", GetTask).Methods(http.MethodGet)
	// r.HandleFunc("/tasks/{id}", UpdateTask).Methods(http.MethodPut)
	// r.HandleFunc("/tasks/{id}", DeleteTask).Methods(http.MethodDelete)
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "Hello World")
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "Hello World")
}
