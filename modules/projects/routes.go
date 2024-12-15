package projects

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammad177/task_management/types"
	"github.com/hammad177/task_management/utils"
)

type Handler struct {
	store types.ProjectStore
}

func NewHandler(store types.ProjectStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", h.createProject).Methods(http.MethodPost)
	r.HandleFunc("/projects", h.getProjects).Methods(http.MethodGet)
	r.HandleFunc("/projects/{project_id}", h.getProjectById).Methods(http.MethodGet)
	r.HandleFunc("/projects/{project_id}", h.updateProject).Methods(http.MethodPatch)
}

func (h *Handler) createProject(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CreateProjectPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.createProjectService(payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

func (h *Handler) getProjects(w http.ResponseWriter, r *http.Request) {
	body, err := h.getProjectsService()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

func (h *Handler) getProjectById(w http.ResponseWriter, r *http.Request) {
	// get the project ID from the URL
	projectID, err := utils.GetURLParams(r, "project_id")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.getProjectByIdService(*projectID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

func (h *Handler) updateProject(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "Hello World")
}
