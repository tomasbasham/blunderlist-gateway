package store

import (
	"context"
	"errors"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"
)

type GetTaskStore interface {
	GetTask(ctx context.Context, id uint) (*entity.Task, error)
}

var ErrTaskNotFound = errors.New("task is not found")

type GetTask struct {
	getTaskStore      GetTaskStore
	listCommentsStore ListCommentsStore
}

func NewGetTask(getTaskStore GetTaskStore, listCommentsStore ListCommentsStore) *GetTask {
	return &GetTask{
		getTaskStore:      getTaskStore,
		listCommentsStore: listCommentsStore,
	}
}

func (s *GetTask) Get(ctx context.Context, id uint) (*entity.Task, error) {
	task, err := s.getTaskStore.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	task.Comments, err = s.listCommentsStore.ListComments(ctx, task.ID)
	if err != nil {
		return nil, err
	}

	return task, nil
}
