package service

import (
	"final_satu/internal/dto"
	"final_satu/internal/entity"
	"final_satu/internal/helper"
	"final_satu/internal/repository"
)

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	var ser service
	ser.repo = repo
	return &ser
}

func (s *service) GetAll() ([]dto.TodosResponseBody, error) {
	var err error
	var getResponses []dto.TodosResponseBody = make([]dto.TodosResponseBody, 0)
	r := s.repo

	err, entities := r.GetAll()
	if err == nil {
		for _, entity := range entities {
			getResponse, err := helper.FromEntityToTodosResponse(entity)
			if err != nil {
				return getResponses, err
			}
			getResponses = append(getResponses, getResponse)
		}
	}

	return getResponses, err

}
func (s *service) GetOne(id int) (dto.TodosResponseBody, error) {
	repo := s.repo
	var err error
	var getResponse dto.TodosResponseBody
	var entity entity.Todos

	err, entity = repo.GetOne(id)
	if err != nil {
		return getResponse, err
	}
	getResponse, err = helper.FromEntityToTodosResponse(entity)
	return getResponse, err

}
func (s *service) Create(newTodo dto.TodosRequestBody) error {
	var err error
	var entity entity.Todos
	repo := s.repo
	entity, err = helper.FromTodosRequestToEntity(newTodo)
	if err != nil {
		return err
	}
	err = repo.Create(entity)
	return err
}
func (s *service) Update(newTodo dto.TodosRequestBody, id int) error {
	var err error
	var entity entity.Todos
	repo := s.repo
	entity, err = helper.FromTodosRequestToEntity(newTodo)
	if err != nil {
		return err
	}
	err = repo.Update(entity, id)
	return err
}
func (s *service) Delete(id int) error {
	var err error
	repo := s.repo

	err = repo.Delete(id)
	return err
}
