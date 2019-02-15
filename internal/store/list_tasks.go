package store

import (
	"context"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"
)

type ListTasksStore interface {
	ListTasks(ctx context.Context) ([]*entity.Task, error)
}

type ListTasks struct {
	listTasksStore    ListTasksStore
	listCommentsStore ListCommentsStore
}

func NewListTasks(listTasksStore ListTasksStore, listCommentsStore ListCommentsStore) *ListTasks {
	return &ListTasks{
		listTasksStore:    listTasksStore,
		listCommentsStore: listCommentsStore,
	}
}

func (s *ListTasks) List(ctx context.Context) ([]*entity.Task, error) {
	tasks, err := s.listTasksStore.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		t.Comments, err = s.listCommentsStore.ListComments(ctx, t.ID)
		if err != nil {
			return nil, err
		}
	}

	return tasks, nil
}
