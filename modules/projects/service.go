package projects

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hammad177/task_management/types"
	"github.com/hammad177/task_management/utils"
)

func (h *Handler) createProjectService(payload types.CreateProjectPayload) (any, error) {
	// validate JSON payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return nil, fmt.Errorf("invalid payload: %v", errors)
	}

	// create project
	if err := h.store.CreateProject(payload); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *Handler) getProjectsService() (any, error) {
	projects, err := h.store.GetProjects()

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (h *Handler) getProjectByIdService(id int) (any, error) {
	project, err := h.store.GetProjectById(id)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (h *Handler) updateProjectService(id int, payload types.UpdateProjectPayload) (any, error) {
	// validate JSON payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return nil, fmt.Errorf("invalid payload: %v", errors)
	}

	// update project
	if err := h.store.UpdateProjectById(id, payload); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *Handler) deleteProjectService(id int) (any, error) {
	err := h.store.DeleteProjectById(id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
