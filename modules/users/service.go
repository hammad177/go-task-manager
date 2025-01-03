package users

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hammad177/task_management/types"
	"github.com/hammad177/task_management/utils"
)

func (h *Handler) createUserService(payload types.CreateUserPayload) (any, error) {
	// validate JSON payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return nil, fmt.Errorf("invalid payload: %v", errors)
	}

	// create user
	if err := h.store.CreateUser(payload); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *Handler) getUserService() (any, error) {
	users, err := h.store.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (h *Handler) getUserByIdService(id int) (any, error) {
	users, err := h.store.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (h *Handler) updateUserService(id int, payload types.UpdateUserPayload) (any, error) {
	// validate JSON payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return nil, fmt.Errorf("invalid payload: %v", errors)
	}

	// update users
	if err := h.store.UpdateUserById(id, payload); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *Handler) deleteUserService(id int) (any, error) {
	err := h.store.DeleteUserById(id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
