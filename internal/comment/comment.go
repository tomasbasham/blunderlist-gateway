package comment

import (
	"time"

	pb "github.com/golang/protobuf/ptypes"
	commentpb "github.com/tomasbasham/blunderlist-comment/blunderlist_comment_v1"
)

// Comment is a type representing a single comment associated with a task.
type Comment struct {
	ID        uint      `jsonapi:"primary,comments"`
	TaskID    uint      `jsonapi:"attr,task_id"`
	Text      string    `jsonapi:"attr,text"`
	Author    string    `jsonapi:"attr,author"`
	CreatedAt time.Time `jsonapi:"attr,created_at,iso8601"`
}

// CommentFromProto returns an internal comment model from the protobuf type.
func CommentFromProto(comment *commentpb.CommentResponse) (*Comment, error) {
	timestamp, err := pb.Timestamp(comment.CreateTime)
	if err != nil {
		return nil, err
	}

	return &Comment{
		ID:        uint(comment.Id),
		Text:      comment.Text,
		CreatedAt: timestamp,
	}, nil
}

// CommentCreateRequestProto returns a comment create request protobuf message
// ready to be sent across the wire.
func CommentCreateRequestProto(comment *Comment) *commentpb.CommentCreateRequest {
	return &commentpb.CommentCreateRequest{
		Text: comment.Text,
	}
}

// CommentUpdateRequestProto returns a comment update request protobuf message
// ready to be sent across the wire.
func CommentUpdateRequestProto(comment *Comment) *commentpb.CommentUpdateRequest {
	return &commentpb.CommentUpdateRequest{
		Id:   uint64(comment.ID),
		Text: comment.Text,
	}
}
