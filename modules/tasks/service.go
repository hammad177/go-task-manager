package tasks

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hammad177/task_management/types"
	"github.com/hammad177/task_management/utils"
)

func (h *Handler) createTaskService(payload types.CreateTaskPayload) (any, error) {
	// validate JSON payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return nil, fmt.Errorf("invalid payload: %v", errors)
	}

	// create task
	if err := h.store.CreateTask(payload); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *Handler) getTaskService(isDetailed bool) (any, error) {
	var tasks any
	var err error

	if isDetailed {
		tasks, err = h.store.GetTasksDetails()
	} else {
		tasks, err = h.store.GetTasks()
	}

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (h *Handler) getTaskByIdService(id int, isDetailed bool) (any, error) {
	var task any
	var err error

	if isDetailed {
		task, err = h.store.GetTaskDetailsById(id)
	} else {
		task, err = h.store.GetTaskById(id)
	}

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (h *Handler) updateTaskService(id int, payload types.UpdateTaskPayload) (any, error) {
	// validate JSON payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return nil, fmt.Errorf("invalid payload: %v", errors)
	}

	// update tasks
	if err := h.store.UpdateTaskById(id, payload); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *Handler) deleteTaskService(id int) (any, error) {
	err := h.store.DeleteTaskById(id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
