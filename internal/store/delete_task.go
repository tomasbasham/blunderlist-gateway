package store

import "context"

type DeleteTaskStore interface {
	DeleteTask(ctx context.Context, id uint) error
}

type DeleteTask struct {
	store DeleteTaskStore
}

func NewDeleteTask(store DeleteTaskStore) *DeleteTask {
	return &DeleteTask{
		store: store,
	}
}

func (s *DeleteTask) Delete(ctx context.Context, id uint) error {
	return s.store.DeleteTask(ctx, id)
}
