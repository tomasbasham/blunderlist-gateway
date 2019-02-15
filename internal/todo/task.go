package todo

import (
	"time"

	pb "github.com/golang/protobuf/ptypes"
	todopb "github.com/tomasbasham/blunderlist-todo/blunderlist_todo_v1"

	"github.com/tomasbasham/blunderlist-gateway/internal/comment"
)

// Task is a type representing a single task.
type Task struct {
	ID        uint               `jsonapi:"primary,tasks"`
	Title     string             `jsonapi:"attr,title"`
	Completed bool               `jsonapi:"attr,completed"`
	CreatedAt time.Time          `jsonapi:"attr,created_at,iso8601"`
	Comments  []*comment.Comment `jsonapi:"relation,comments"`
}

// TaskFromProto returns an internal task model from the protobuf type.
func TaskFromProto(task *todopb.TaskResponse) (*Task, error) {
	timestamp, err := pb.Timestamp(task.CreateTime)
	if err != nil {
		return nil, err
	}

	return &Task{
		ID:        uint(task.Id),
		Title:     task.Title,
		Completed: task.Completed,
		CreatedAt: timestamp,
	}, nil
}

// TaskCreateRequestProto returns a task create request protobuf message ready
// to be sent across the wire.
func TaskCreateRequestProto(task *Task) *todopb.TaskCreateRequest {
	return &todopb.TaskCreateRequest{
		Title:     task.Title,
		Completed: task.Completed,
	}
}

// TaskUpdateRequestProto returns a task update request protobuf message ready
// to be sent across the wire.
func TaskUpdateRequestProto(task *Task) *todopb.TaskUpdateRequest {
	return &todopb.TaskUpdateRequest{
		Id:        uint64(task.ID),
		Title:     task.Title,
		Completed: task.Completed,
	}
}
