package store

import (
	"context"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"
)

type ListCommentsStore interface {
	ListComments(ctx context.Context, pID uint) ([]*entity.Comment, error)
}

type ListComments struct {
	store ListCommentsStore
}

func NewListComments(store ListCommentsStore) *ListComments {
	return &ListComments{
		store: store,
	}
}

func (s *ListComments) List(ctx context.Context, pID uint) ([]*entity.Comment, error) {
	return s.store.ListComments(ctx, pID)
}
