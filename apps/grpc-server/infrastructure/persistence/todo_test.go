package persistence

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taguch1/try-bff/apps/grpc-server/infrastructure/mysql"
)

func TestHoge(t *testing.T) {
	ctx := context.Background()
	mysqlConfig, err := mysql.NewConf("../../config/mysql.json")
	assert.Nil(t, err)

	db, err := mysql.Open(mysqlConfig)
	assert.Nil(t, err)

	todoRepo := NewTodo(db)

	title := "TitleX"
	todo, err := todoRepo.Save(ctx, title)
	assert.Nil(t, err)
	assert.Equal(t, title, todo.Title)

	id := "ID1"
	todo, err = todoRepo.Get(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, "ID1", todo.ID)
	assert.Equal(t, "TitleA", todo.Title)

	todos, err := todoRepo.List(ctx, 0, 10)
	assert.Nil(t, err)
	assert.Equal(t, 5, len(todos))

}
