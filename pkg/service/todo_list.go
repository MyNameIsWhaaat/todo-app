package service

import (
	"github.com/MyNameIsWhaaat/todo-app"
	"github.com/MyNameIsWhaaat/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func newTodoListService(repo repository.TodoList) *TodoListService{
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error){
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error){
	return s.repo.GetAll(userId)
}