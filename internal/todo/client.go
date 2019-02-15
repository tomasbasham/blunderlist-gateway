package todo

import (
	"context"
	"io"
	"log"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"

	pb "github.com/golang/protobuf/ptypes"
	todopb "github.com/tomasbasham/blunderlist-todo/blunderlist_todo_v1"
)

type Client struct {
	logger *log.Logger
	client todopb.TodoClient
}

func NewClient(logger *log.Logger, client todopb.TodoClient) *Client {
	return &Client{
		logger: logger,
		client: client,
	}
}

func (c *Client) ListTasks(ctx context.Context) ([]*entity.Task, error) {
	stream, err := c.client.ListTasks(ctx, &todopb.TaskListRequest{})
	if err != nil {
		return nil, err
	}

	return taskList(stream)
}

func (c *Client) GetTask(ctx context.Context, id uint) (*entity.Task, error) {
	task, err := c.client.GetTask(ctx, &todopb.TaskQuery{Id: uint64(id)})
	if err != nil {
		return nil, err
	}

	return taskFromProto(task)
}

func (c *Client) CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	t, err := c.client.CreateTask(ctx, taskCreateRequestProto(task))
	if err != nil {
		return nil, err
	}

	return taskFromProto(t)
}

func (c *Client) UpdateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	t, err := c.client.UpdateTask(ctx, taskUpdateRequestProto(task))
	if err != nil {
		return nil, err
	}

	return taskFromProto(t)
}

func (c *Client) DeleteTask(ctx context.Context, id uint) error {
	if _, err := c.client.DeleteTask(ctx, &todopb.TaskQuery{Id: uint64(id)}); err != nil {
		return err
	}

	return nil
}

func taskList(stream todopb.Todo_ListTasksClient) ([]*entity.Task, error) {
	var tasks []*entity.Task

	for {
		t, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		task, err := taskFromProto(t)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// taskFromProto returns an internal task model from the protobuf type.
func taskFromProto(task *todopb.TaskResponse) (*entity.Task, error) {
	timestamp, err := pb.Timestamp(task.CreateTime)
	if err != nil {
		return nil, err
	}

	return &entity.Task{
		ID:        uint(task.Id),
		Title:     task.Title,
		Completed: task.Completed,
		CreatedAt: timestamp,
	}, nil
}

// taskCreateRequestProto returns a task create request protobuf message ready
// to be sent across the wire.
func taskCreateRequestProto(task *entity.Task) *todopb.TaskCreateRequest {
	return &todopb.TaskCreateRequest{
		Title:     task.Title,
		Completed: task.Completed,
	}
}

// taskUpdateRequestProto returns a task update request protobuf message ready
// to be sent across the wire.
func taskUpdateRequestProto(task *entity.Task) *todopb.TaskUpdateRequest {
	return &todopb.TaskUpdateRequest{
		Id:        uint64(task.ID),
		Title:     task.Title,
		Completed: task.Completed,
	}
}
