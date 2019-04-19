package repository

import (
	"context"

	"github.com/taguch1/try-bff/apps/grpc-server/domain/model"
)

// Todo service interface
type Todo interface {
	NextID(ctx context.Context) string
	Save(ctx context.Context, id, title string) (*model.Todo, error)
	Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context, offset, limit int64) ([]*model.Todo, error)
	Update(ctx context.Context, id, title string) (*model.Todo, error)
	Delete(ctx context.Context, id string) error
}
