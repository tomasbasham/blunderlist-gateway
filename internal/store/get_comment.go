package store

import (
	"context"
	"errors"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"
)

type GetCommentStore interface {
	GetComment(ctx context.Context, id uint) (*entity.Comment, error)
}

var ErrCommentNotFound = errors.New("comment is not found")

type GetComment struct {
	store GetCommentStore
}

func NewGetComment(store GetCommentStore) *GetComment {
	return &GetComment{
		store: store,
	}
}

func (s *GetComment) Get(ctx context.Context, id uint) (*entity.Comment, error) {
	return s.store.GetComment(ctx, id)
}
