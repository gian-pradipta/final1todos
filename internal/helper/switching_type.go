package helper

import (
	"final_satu/internal/dto"
	"final_satu/internal/entity"
)

func FromEntityToTodosResponse(todos entity.Todos) (dto.TodosResponseBody, error) {
	var err error
	var getResponse dto.TodosResponseBody

	getResponse.Nama = todos.Nama
	getResponse.Deskripsi = todos.Deskripsi
	getResponse.TanggalMulai = todos.TanggalMulai
	getResponse.TanggalSelesai = todos.TanggalSelesai

	return getResponse, err
}

func FromTodosRequestToEntity(createRequest dto.TodosRequestBody) (entity.Todos, error) {
	var err error
	var todo entity.Todos

	todo.Nama = createRequest.Nama
	todo.TanggalMulai = createRequest.TanggalMulai
	todo.TanggalSelesai = createRequest.TanggalSelesai
	todo.Deskripsi = createRequest.Deskripsi

	return todo, err
}
