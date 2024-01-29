package service

import (
	"example.com/tma/controller/request"
	"example.com/tma/controller/response"
)

type TaskService interface {
	Add(add request.TaskAdd) error
	GetTask(string) (*response.TaskGet, error)
}
