package service

import (
	"final_satu/internal/dto"
)

type Service interface {
	GetAll() ([]dto.TodosResponseBody, error)
	GetOne(id int) (dto.TodosResponseBody, error)
	Create(newTodo dto.TodosRequestBody) error
	Update(newTodo dto.TodosRequestBody, id int) error
	Delete(id int) error
}
