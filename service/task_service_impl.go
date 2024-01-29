package service

import (
	"errors"
	"example.com/tma/controller/request"
	"example.com/tma/controller/response"
	"example.com/tma/entity"
	"example.com/tma/repository"
)

type TaskServiceImpl struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return TaskServiceImpl{
		taskRepository: taskRepository,
	}
}

func (tsk TaskServiceImpl) Add(taskAddReq request.TaskAdd) error {
	task := entity.Task{Title: taskAddReq.Title, Description: taskAddReq.Description}
	err := tsk.taskRepository.CreateTask(task)

	if err != nil {
		return errors.New("failed to save task to db")
	}

	return nil
}

func (tsk TaskServiceImpl) GetTask(id string) (*response.TaskGet, error) {
	taskResult, err := tsk.taskRepository.GetTask(id)
	if err != nil {
		return nil, err
	}
	taskResponse := response.TaskGet{
		Title:       taskResult.Title,
		Description: taskResult.Description,
	}

	return &taskResponse, nil
}
