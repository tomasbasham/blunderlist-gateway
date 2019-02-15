package store

import (
	"context"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"
)

type UpdateCommentStore interface {
	UpdateComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)
}

type UpdateComment struct {
	store UpdateCommentStore
}

func NewUpdateComment(store UpdateCommentStore) *UpdateComment {
	return &UpdateComment{
		store: store,
	}
}

func (s *UpdateComment) Update(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	return s.store.UpdateComment(ctx, comment)
}
