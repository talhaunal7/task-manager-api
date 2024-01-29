package main

import (
	"example.com/tma/controller"
	"example.com/tma/entity"
	"example.com/tma/repository"
	"example.com/tma/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	server         *gin.Engine
	taskRepository repository.TaskRepository
	taskService    service.TaskService
	taskController controller.TaskController
	db             *gorm.DB
	err            error
)

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&entity.Task{})

	if err != nil {
		log.Fatal(err.Error())
	}

	taskRepository = repository.NewTaskRepository(db)
	taskService = service.NewTaskService(taskRepository)
	taskController = controller.NewTaskController(taskService)

	server = gin.Default()

}

func main() {
	basepath := server.Group("/v1")
	taskController.RegisterRoutes(basepath)
	log.Fatal(server.Run(":8080"))
}
