package main

import (
	"log"

	"https://github.com/golangtestcases/task/internal/controller/task"
	"https://github.com/golangtestcases/task/internal/repository"
	"https://github.com/golangtestcases/task/internal/server"
	"https://github.com/golangtestcases/task/internal/usecase"
)

func main() {
	taskRepo := repository.NewInMemoryTaskRepository()
	taskUseCase := usecase.NewTaskUseCase(taskRepo)
	taskHandler := task.NewHandler(*taskUseCase)

	srv := server.NewServer(taskHandler)
	log.Fatal(srv.Run(":8080"))
}
