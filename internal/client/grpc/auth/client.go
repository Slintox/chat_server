package auth

import (
	"context"
	"github.com/Slintox/chat_server/internal/client/grpc/converter"
	"github.com/Slintox/chat_server/internal/model"
	"github.com/Slintox/user-service/pkg/user_v1"
	"google.golang.org/grpc"
)

type Client interface {
	Get(ctx context.Context, username string) (*model.User, error)
}

type client struct {
	noteClient user_v1.UserV1Client
}

func NewClient(conn *grpc.ClientConn) Client {
	return &client{
		noteClient: user_v1.NewUserV1Client(conn),
	}
}

func (c *client) Get(ctx context.Context, username string) (*model.User, error) {
	user, err := c.noteClient.Get(ctx, converter.ToGetUserRequest(username))
	if err != nil {
		return nil, err
	}

	return converter.FromGetUserResponse(user), nil
}
