package persistence

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/taguch1/try-bff/apps/grpc-server/domain/model"
	"github.com/taguch1/try-bff/apps/grpc-server/domain/repository"
)

type todoRepoImpl struct {
	db *sqlx.DB
}

// NewTodo constructor
func NewTodo(db *sqlx.DB) repository.Todo {
	return &todoRepoImpl{db}
}

func (r *todoRepoImpl) NextID(ctx context.Context) string {
	return xid.New().String()
}

func (r *todoRepoImpl) Save(ctx context.Context, id, title string) (*model.Todo, error) {
	query := `INSERT INTO  rdbms.todo (id,title) VALUES (:id,:title)`
	_, err := r.db.NamedExec(query,
		map[string]interface{}{
			"id":    id,
			"title": title,
		})
	if err != nil {
		return nil, err
	}
	return &model.Todo{ID: id, Title: title}, nil
}

func (r *todoRepoImpl) Get(ctx context.Context, id string) (*model.Todo, error) {
	query := `
		SELECT
			id,
			title
		FROM
			rdbms.todo
		WHERE
			id = ?
		LIMIT 1`

	var todo model.Todo

	if err := r.db.Get(&todo, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepoImpl) List(ctx context.Context, offset, limit int64) ([]*model.Todo, error) {
	query := `
		SELECT
			id,
			title
		FROM
			rdbms.todo
		LIMIT ?, ?`

	var todos []*model.Todo

	if err := r.db.Select(&todos, query, offset, limit); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return todos, nil

}

func (r *todoRepoImpl) Update(ctx context.Context, id, title string) (*model.Todo, error) {
	query := `UPDATE rdbms.todo SET title= :title WHERE id = :id LIMIT 1`
	result, err := r.db.NamedExec(query,
		map[string]interface{}{
			"id":    id,
			"title": title,
		})
	if err != nil {
		return nil, err
	}

	if affectedRows, _ := result.RowsAffected(); affectedRows == 0 {
		return nil, nil
	}
	todo, err := r.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepoImpl) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM rdbms.todo WHERE id = :id LIMIT 1`
	_, err := r.db.NamedExec(query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return err
	}
	return nil
}
