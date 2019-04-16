package service

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/taguch1/try-bff/apps/bff-grpc/domain/model"
)

// TodoService mock
type TodoService struct {
	mock.Mock
}

// Save mock
func (m *TodoService) Save(ctx context.Context, title string) (*model.Todo, error) {
	args := m.Called(ctx, title)
	return args.Get(0).(*model.Todo), args.Error(1)
}

// Get mock
func (m *TodoService) Get(ctx context.Context, id string) (*model.Todo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

// List mock
func (m *TodoService) List(ctx context.Context, offset, limit int64) ([]*model.Todo, error) {
	args := m.Called(ctx, offset, limit)
	return args.Get(0).([]*model.Todo), args.Error(1)
}

// Update mock
func (m *TodoService) Update(ctx context.Context, id, title string) (*model.Todo, error) {
	args := m.Called(ctx, id, title)
	return args.Get(0).(*model.Todo), args.Error(1)
}

// Delete mock
func (m *TodoService) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
