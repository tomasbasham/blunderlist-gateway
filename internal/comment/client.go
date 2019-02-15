package comment

import (
	"context"
	"io"
	"log"

	"github.com/tomasbasham/blunderlist-gateway/internal/entity"

	pb "github.com/golang/protobuf/ptypes"
	commentpb "github.com/tomasbasham/blunderlist-comment/blunderlist_comment_v1"
)

type Client struct {
	logger *log.Logger
	client commentpb.CommentClient
}

func NewClient(logger *log.Logger, client commentpb.CommentClient) *Client {
	return &Client{
		logger: logger,
		client: client,
	}
}

func (c *Client) ListComments(ctx context.Context, pID uint) ([]*entity.Comment, error) {
	stream, err := c.client.ListComments(ctx, &commentpb.CommentListRequest{ParentId: uint64(pID)})
	if err != nil {
		return nil, err
	}

	return commentList(stream)
}

func (c *Client) GetComment(ctx context.Context, id uint) (*entity.Comment, error) {
	comment, err := c.client.GetComment(ctx, &commentpb.CommentQuery{Id: uint64(id)})
	if err != nil {
		return nil, err
	}

	return commentFromProto(comment)
}

func (c *Client) CreateComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	cc, err := c.client.CreateComment(ctx, commentCreateRequestProto(comment))
	if err != nil {
		return nil, err
	}

	return commentFromProto(cc)
}

func (c *Client) UpdateComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	cc, err := c.client.UpdateComment(ctx, commentUpdateRequestProto(comment))
	if err != nil {
		return nil, err
	}

	return commentFromProto(cc)
}

func (c *Client) DeleteComment(ctx context.Context, id uint) error {
	if _, err := c.client.DeleteComment(ctx, &commentpb.CommentQuery{Id: uint64(id)}); err != nil {
		return err
	}

	return nil
}

func commentList(stream commentpb.Comment_ListCommentsClient) ([]*entity.Comment, error) {
	var comments []*entity.Comment

	for {
		c, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		comment, err := commentFromProto(c)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

// commentFromProto returns an internal comment model from the protobuf type.
func commentFromProto(comment *commentpb.CommentResponse) (*entity.Comment, error) {
	timestamp, err := pb.Timestamp(comment.CreateTime)
	if err != nil {
		return nil, err
	}

	return &entity.Comment{
		ID:        uint(comment.Id),
		Text:      comment.Text,
		CreatedAt: timestamp,
	}, nil
}

// commentCreateRequestProto returns a comment create request protobuf message
// ready to be sent across the wire.
func commentCreateRequestProto(comment *entity.Comment) *commentpb.CommentCreateRequest {
	return &commentpb.CommentCreateRequest{
		Text: comment.Text,
	}
}

// commentUpdateRequestProto returns a comment update request protobuf message
// ready to be sent across the wire.
func commentUpdateRequestProto(comment *entity.Comment) *commentpb.CommentUpdateRequest {
	return &commentpb.CommentUpdateRequest{
		Id:   uint64(comment.ID),
		Text: comment.Text,
	}
}
