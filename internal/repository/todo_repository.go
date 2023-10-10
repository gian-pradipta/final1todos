package repository

import "final_satu/internal/entity"

type Repository interface {
	GetAll() (error, []entity.Todos)
	GetOne(id uint) (error, entity.Todos)
	Create(newTodo entity.Todos) error
	Update(newTodo entity.Todos, id uint) error
	Delete(id uint) error
}
