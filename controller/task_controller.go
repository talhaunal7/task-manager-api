package controller

import (
	"example.com/tma/controller/request"
	"example.com/tma/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskController struct {
	taskService service.TaskService
}

func NewTaskController(taskService service.TaskService) TaskController {
	return TaskController{
		taskService: taskService,
	}
}

func (tsk TaskController) Add(ctx *gin.Context) {
	var task request.TaskAdd
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tsk.taskService.Add(task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (tsk TaskController) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := tsk.taskService.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)

}

func (tsk TaskController) RegisterRoutes(rg *gin.RouterGroup) {
	taskRoute := rg.Group("/tasks")
	taskRoute.POST("/", tsk.Add)
	taskRoute.GET("/:id", tsk.GetTask)
}
