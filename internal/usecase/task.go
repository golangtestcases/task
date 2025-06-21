package usecase

import (
	"context"
	"math/rand"
	"time"

	"https://github.com/golangtestcases/task/internal/entity"
	"https://github.com/golangtestcases/task/internal/repository"
)

type TaskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) *TaskUseCase {
	rand.Seed(time.Now().UnixNano())
	return &TaskUseCase{taskRepo: repo}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context) (*entity.Task, error) {
	now := time.Now()
	task := &entity.Task{
		ID:        entity.GenerateID(),
		Status:    entity.StatusPending,
		CreatedAt: now,
	}

	if err := uc.taskRepo.Create(task); err != nil {
		return nil, err
	}

	go uc.processTask(task.ID)

	return task, nil
}

func (uc *TaskUseCase) GetTask(ctx context.Context, id string) (*entity.Task, error) {
	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if task.StartedAt != nil {
		if task.CompletedAt != nil {
			task.DurationMs = task.CompletedAt.Sub(*task.StartedAt).Milliseconds()
		} else {
			task.DurationMs = time.Since(*task.StartedAt).Milliseconds()
		}
	}

	return task, nil
}

func (uc *TaskUseCase) DeleteTask(ctx context.Context, id string) error {
	return uc.taskRepo.Delete(id)
}

func (uc *TaskUseCase) processTask(id string) {
	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return
	}

	startedAt := time.Now()
	task.Status = entity.StatusProcessing
	task.StartedAt = &startedAt
	if err := uc.taskRepo.Update(task); err != nil {
		return
	}

	duration := 3*time.Minute + time.Duration(rand.Int63n(2*60))*time.Second
	time.Sleep(duration)

	completedAt := time.Now()
	result := "task completed successfully"
	task.Status = entity.StatusCompleted
	task.CompletedAt = &completedAt
	task.Result = &result
	_ = uc.taskRepo.Update(task)
}
