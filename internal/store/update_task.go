package store

import (
	"context"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"
)

type UpdateTaskStore interface {
	UpdateTask(ctx context.Context, task *entity.Task) (*entity.Task, error)
}

type UpdateTask struct {
	store UpdateTaskStore
}

func NewUpdateTask(store UpdateTaskStore) *UpdateTask {
	return &UpdateTask{
		store: store,
	}
}

func (s *UpdateTask) Update(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	return s.store.UpdateTask(ctx, task)
}
