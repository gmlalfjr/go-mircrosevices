package services

import (
	"time"

	"github.com/go-microservices/todo-list/domain"
	errors "github.com/go-microservices/todo-list/utils"
)

func AddTodoService(t *domain.Todo, userId string) (*domain.Todo, *errors.RestErr) {
	date := time.Now().UTC()
	t.CreatedAt = date
	t.ModifiedAt = date
	t.Done = 0
	t.UserId = userId

	err := t.TodoDao()
	if err != nil {
		return nil, errors.NewBadRequest("something error")
	}

	return t, nil
}

func UpdateTodoService(t *domain.Todo, idTodo int64, isPartial bool) (*domain.Todo, *errors.RestErr) {
	current, err := GetTodoDetailService(idTodo)
	if err != nil {
		return nil, errors.NewBadRequest(err.Message)
	}
	date := time.Now().UTC()
	t.Id = idTodo
	t.ModifiedAt = date
	t.CreatedAt = current.CreatedAt
	if isPartial {
		if t.Description == "" {
			t.Description = current.Description
		}
		if t.Done == 0 {
			t.Done = 0
		} else if t.Done == 1 {
			t.Done = 1
		}
	}

	if err := t.TodoDaoUpdate(); err != nil {
		return nil, errors.NewBadRequest("something error")
	}

	return t, nil
}

func GetTodoDetailService(todoId int64) (*domain.Todo, *errors.RestErr) {
	result := &domain.Todo{Id: todoId}
	if err := result.GetTodoDetailDao(); err != nil {
		return nil, errors.NewBadRequest(err.Message)
	}

	return result, nil
}

func GetTodoService(t *domain.Todo) ([]domain.Todo, *errors.RestErr) {
	return t.GetTodoDao()
}

func DeleteTodoService(todoId int64) *errors.RestErr {
	result := &domain.Todo{Id: todoId}
	result.DeleteTodoDao()

	return nil
}
