package repository

import (
	"errors"
	"sync"

	"github.com/golangtestcases/task/internal/entity"
)

type inMemoryTaskRepository struct {
	tasks map[string]*entity.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepository() *inMemoryTaskRepository {
	return &inMemoryTaskRepository{
		tasks: make(map[string]*entity.Task),
	}
}

func (r *inMemoryTaskRepository) Create(task *entity.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.ID] = task
	return nil
}

func (r *inMemoryTaskRepository) GetByID(id string) (*entity.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}

	return task, nil
}

func (r *inMemoryTaskRepository) Update(task *entity.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return errors.New("task not found")
	}

	r.tasks[task.ID] = task
	return nil
}

func (r *inMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(r.tasks, id)
	return nil
}
