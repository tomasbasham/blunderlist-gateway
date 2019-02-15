package store

import (
	"context"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"
)

type CreateCommentStore interface {
	CreateComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)
}

type CreateComment struct {
	store CreateCommentStore
}

func NewCreateComment(store CreateCommentStore) *CreateComment {
	return &CreateComment{
		store: store,
	}
}

func (s *CreateComment) Create(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	return s.store.CreateComment(ctx, comment)
}
