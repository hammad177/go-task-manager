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
	r.HandleFunc("/tasks", h.createTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks", h.getTasks).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{task_id}", h.getTaskById).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{task_id}", h.updateTask).Methods(http.MethodPatch)
	r.HandleFunc("/tasks/{task_id}", h.deleteTask).Methods(http.MethodDelete)
}

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CreateTaskPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.createTaskService(payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, body)
}

func (h *Handler) getTasks(w http.ResponseWriter, r *http.Request) {
	body, err := h.getTaskService()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

func (h *Handler) getTaskById(w http.ResponseWriter, r *http.Request) {
	// get the task ID from the URL
	taskID, err := utils.GetURLParams(r, "task_id")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.getTaskByIdService(*taskID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.UpdateTaskPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// get the task ID from the URL
	taskID, err := utils.GetURLParams(r, "task_id")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.updateTaskService(*taskID, payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	// get the task ID from the URL
	taskID, err := utils.GetURLParams(r, "task_id")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.deleteTaskService(*taskID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}
