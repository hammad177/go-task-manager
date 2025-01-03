package users

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammad177/task_management/types"
	"github.com/hammad177/task_management/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", h.createUser).Methods(http.MethodPost)
	r.HandleFunc("/users", h.getUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}", h.getUserById).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}", h.updateUser).Methods(http.MethodPatch)
	r.HandleFunc("/users/{user_id}", h.deleteUser).Methods(http.MethodDelete)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CreateUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.createUserService(payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, body)
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	body, err := h.getUserService()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request) {
	// get the user ID from the URL
	userID, err := utils.GetURLParams(r, "user_id")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.getUserByIdService(*userID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.UpdateUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// get the user ID from the URL
	userID, err := utils.GetURLParams(r, "user_id")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.updateUserService(*userID, payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	// get the user ID from the URL
	userID, err := utils.GetURLParams(r, "user_id")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	body, err := h.deleteUserService(*userID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, body)
}
