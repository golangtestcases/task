package main

import (
	"log"

	"github.com/golangtestcases/task/internal/controller/task"
	"github.com/golangtestcases/task/internal/repository"
	"github.com/golangtestcases/task/internal/server"
	"github.com/golangtestcases/task/internal/usecase"
)

func main() {
	taskRepo := repository.NewInMemoryTaskRepository()
	taskUseCase := usecase.NewTaskUseCase(taskRepo)
	taskHandler := task.NewHandler(*taskUseCase)

	srv := server.NewServer(taskHandler)
	log.Fatal(srv.Run(":8080"))
}
