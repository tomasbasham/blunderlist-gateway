package store

import (
	"context"
	"io"

	commentpb "github.com/tomasbasham/blunderlist-comment/blunderlist_comment_v1"
	todopb "github.com/tomasbasham/blunderlist-todo/blunderlist_todo_v1"

	"github.com/tomasbasham/blunderlist-gateway/internal/comment"
	"github.com/tomasbasham/blunderlist-gateway/internal/todo"
)

// Store
type Store struct {
	todo    todopb.TodoClient
	comment commentpb.CommentClient
}

// New
func New(todo todopb.TodoClient, comment commentpb.CommentClient) *Store {
	return &Store{
		todo:    todo,
		comment: comment,
	}
}

func (s *Store) ListTasks(ctx context.Context) ([]*todo.Task, error) {
	stream, err := s.todo.ListTasks(ctx, &todopb.TaskListRequest{})
	if err != nil {
		return nil, err
	}

	tasks, err := s.taskList(stream)
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		t.Comments, err = s.ListComments(ctx, t.ID)
		if err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

func (s *Store) taskList(stream todopb.Todo_ListTasksClient) ([]*todo.Task, error) {
	var tasks []*todo.Task

	for {
		t, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		task, err := todo.TaskFromProto(t)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Store) GetTask(ctx context.Context, id uint) (*todo.Task, error) {
	t, err := s.todo.GetTask(ctx, &todopb.TaskQuery{Id: uint64(id)})
	if err != nil {
		return nil, err
	}

	task, err := todo.TaskFromProto(t)
	if err != nil {
		return nil, err
	}

	task.Comments, err = s.ListComments(ctx, task.ID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Store) CreateTask(ctx context.Context, task *todo.Task) (*todo.Task, error) {
	t, err := s.todo.CreateTask(ctx, todo.TaskCreateRequestProto(task))
	if err != nil {
		return nil, err
	}

	return todo.TaskFromProto(t)
}

func (s *Store) UpdateTask(ctx context.Context, task *todo.Task) (*todo.Task, error) {
	t, err := s.todo.UpdateTask(ctx, todo.TaskUpdateRequestProto(task))
	if err != nil {
		return nil, err
	}

	return todo.TaskFromProto(t)
}

func (s *Store) DeleteTask(ctx context.Context, id uint) error {
	if _, err := s.todo.DeleteTask(ctx, &todopb.TaskQuery{Id: uint64(id)}); err != nil {
		return err
	}

	return nil
}

func (s *Store) ListComments(ctx context.Context, pID uint) ([]*comment.Comment, error) {
	stream, err := s.comment.ListComments(ctx, &commentpb.CommentListRequest{ParentId: uint64(pID)})
	if err != nil {
		return nil, err
	}

	return s.commentList(stream)
}

func (s *Store) commentList(stream commentpb.Comment_ListCommentsClient) ([]*comment.Comment, error) {
	var comments []*comment.Comment

	for {
		c, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		comment, err := comment.CommentFromProto(c)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (s *Store) GetComment(ctx context.Context, id uint) (*comment.Comment, error) {
	c, err := s.comment.GetComment(ctx, &commentpb.CommentQuery{Id: uint64(id)})
	if err != nil {
		return nil, err
	}

	return comment.CommentFromProto(c)
}

func (s *Store) CreateComment(ctx context.Context, comm *comment.Comment) (*comment.Comment, error) {
	c, err := s.comment.CreateComment(ctx, comment.CommentCreateRequestProto(comm))
	if err != nil {
		return nil, err
	}

	return comment.CommentFromProto(c)
}

func (s *Store) UpdateComment(ctx context.Context, comm *comment.Comment) (*comment.Comment, error) {
	c, err := s.comment.UpdateComment(ctx, comment.CommentUpdateRequestProto(comm))
	if err != nil {
		return nil, err
	}

	return comment.CommentFromProto(c)
}

func (s *Store) DeleteComment(ctx context.Context, id uint) error {
	if _, err := s.comment.DeleteComment(ctx, &commentpb.CommentQuery{Id: uint64(id)}); err != nil {
		return err
	}

	return nil
}
