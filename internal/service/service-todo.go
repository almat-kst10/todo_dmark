package service

import (
	"context"
	"todo_dmark/internal/models"
	"todo_dmark/internal/repository"
)

type ITodoService interface {
	List(ctx context.Context) ([]models.TodoItem, error)
	Add(ctx context.Context, todo *models.TodoItem) (error)
	Detail(ctx context.Context, id int) (models.TodoItem, error)
	Update(ctx context.Context, item models.TodoItem) (bool, error)
	Delete(ctx context.Context, id int) error
}

type TodoService struct {
	todoRepo repository.ITodoRepo
}

func NewTodoService(todoRepo repository.ITodoRepo) ITodoService {
	return &TodoService{
		todoRepo: todoRepo,
	}
}

func (s *TodoService) List(ctx context.Context) ([]models.TodoItem, error) {
	return s.todoRepo.List(ctx)
}

func (s *TodoService) Add(ctx context.Context, todo *models.TodoItem) (error) {
	return s.todoRepo.Add(ctx, todo)
}

func (s *TodoService) Detail(ctx context.Context, id int) (models.TodoItem, error) {
	return s.todoRepo.Detail(ctx, id)
}

func (s *TodoService) Update(ctx context.Context, item models.TodoItem) (bool, error) {
	return s.todoRepo.Update(ctx, item)
}

func (s *TodoService) Delete(ctx context.Context, id int) error {
	return s.todoRepo.Delete(ctx, id)
}
