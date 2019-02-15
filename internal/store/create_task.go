package store

import (
	"context"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"
)

type CreateTaskStore interface {
	CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error)
}

type CreateTask struct {
	store CreateTaskStore
}

func NewCreateTask(store CreateTaskStore) *CreateTask {
	return &CreateTask{
		store: store,
	}
}

func (s *CreateTask) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	return s.store.CreateTask(ctx, task)
}
