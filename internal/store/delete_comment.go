package store

import "context"

type DeleteCommentStore interface {
	DeleteComment(ctx context.Context, id uint) error
}

type DeleteComment struct {
	store DeleteCommentStore
}

func NewDeleteComment(store DeleteCommentStore) *DeleteComment {
	return &DeleteComment{
		store: store,
	}
}

func (s *DeleteComment) Delete(ctx context.Context, id uint) error {
	return s.store.DeleteComment(ctx, id)
}
