package repository

import (
	"errors"
	"example.com/tma/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task entity.Task) error
	GetTask(id string) (*entity.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return taskRepository{db: db}
}

func (taskRepo taskRepository) CreateTask(task entity.Task) error {
	result := taskRepo.db.Create(&task)
	return result.Error
}

func (taskRepo taskRepository) GetTask(id string) (*entity.Task, error) {
	var task entity.Task
	result := taskRepo.db.First(&task, id)

	if result.Error != nil {
		return nil, errors.New("could not found")
	}
	return &task, nil

}
