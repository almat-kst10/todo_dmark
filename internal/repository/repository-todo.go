package repository

import (
	"context"
	"database/sql"
	"todo_dmark/internal/models"
)

type ITodoRepo interface {
	List(ctx context.Context) ([]models.TodoItem, error)
	Add(ctx context.Context, todo *models.TodoItem) (error)
	Detail(ctx context.Context, id int) (models.TodoItem, error)
	Update(ctx context.Context, item models.TodoItem) (bool, error)
	Delete(ctx context.Context, id int) error
}

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) ITodoRepo {
	return &TodoRepo{
		db: db,
	}
}

func (r *TodoRepo) List(ctx context.Context) ([]models.TodoItem, error) {

	query := `SELECT id, title, time_at, is_done FROM todos`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.TodoItem

	for rows.Next() {
		var item models.TodoItem
		err := rows.Scan(&item.Id, &item.Title, &item.TimeAt, &item.IsDone)
		if err != nil {
			return nil, err
		}

		list = append(list, item)
	}

	return list, nil
}

func (r *TodoRepo) Add(ctx context.Context, todo *models.TodoItem) (error) {
	query := "INSERT INTO todos(title, is_done, time_at) VALUES($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, todo.Title, todo.IsDone, todo.TimeAt)
	return err
}

func (r *TodoRepo) Detail(ctx context.Context, id int) (models.TodoItem, error) {
	var todo models.TodoItem

	query := `SELECT id, title, time_at, is_done FROM todos WHERE id=$1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&todo.Id, &todo.Title, &todo.TimeAt, &todo.IsDone)

	return todo, err
}

func (r *TodoRepo) Update(ctx context.Context, item models.TodoItem) (bool, error) {
	query := "UPDATE todos SET title = $1, is_done = $2 WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, item.Title, item.IsDone, item.Id)
	return item.IsDone, err
}

func (r *TodoRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
