package controller

import (
	"example.com/tma/service"
)

type TaskController struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) TaskController {
	return TaskController{
		TaskService: taskService,
	}
}
