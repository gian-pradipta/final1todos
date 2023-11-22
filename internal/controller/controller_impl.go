package controller

import (
	"final_satu/internal/dto"
	"final_satu/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.Service
}

func New(service service.Service) Controller {
	var c controller

	c.service = service

	return &c
}

// ListAccounts godoc
// @Summary      One todos
// @Description  get one todos
// @Tags         todos
// @Produce      json
// @param id path int true "ID of todos to be presented"
// @Success      200  {object} dto.TodosResponseBody
// @Failure      400  {object}  dto.HTTPError
// @Failure      404  {object}  dto.HTTPError
// @Router       /todos/:id [get]
func (ctr *controller) GetOne(c *gin.Context) {
	var err error
	var response dto.TodosResponseBody
	var httpError dto.HTTPError

	id := c.Param("id")
	validId, err := validateId(id)
	if err != nil {
		httpError = dto.NewHTTPError(http.StatusBadRequest, "Invalid ID")
		goto ERROR_HANDLER
	}

	response, err = ctr.service.GetOne(validId)
	if err != nil {
		httpError = dto.NewHTTPError(http.StatusNotFound, "Failed to get Data")
		goto ERROR_HANDLER
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
	return

ERROR_HANDLER:
	c.JSON(httpError.ErrCode, httpError)
	return
}

// ListAccounts godoc
// @Summary      List todoses
// @Description  get all todoses
// @Tags         todos
// @Produce      json
// @Success      200  {array} dto.TodosResponseBody
// @Failure      400  {object}  dto.HTTPError
// @Failure      404  {object}  dto.HTTPError
// @Router       /todos/ [get]
func (ctr *controller) GetAll(c *gin.Context) {
	s := ctr.service
	var err error
	var httpCode int = http.StatusBadRequest
	var httpError dto.HTTPError
	var responses []dto.TodosResponseBody
	responses, err = s.GetAll()
	if err != nil {
		httpError = dto.NewHTTPError(httpCode, "Failed to get data")
		goto ERROR_HANDLER
	}
	c.JSON(http.StatusOK, gin.H{
		"data": responses,
	})
	return

ERROR_HANDLER:
	if err != nil {
		c.JSON(httpError.ErrCode, httpError)
	}
	return
}

// ListAccounts godoc
// @Summary      Create new Todos
// @Description  Create New Todos based on valid json request
// @Tags         todos
// @Accept       json
// @Success      200  {object} dto.TodosResponseBody
// @Failure      400  {object}  dto.HTTPError
// @Router       /todos/ [post]
func (ctr *controller) Create(c *gin.Context) {
	var err error
	var httpCode int = http.StatusBadRequest
	var httpError dto.HTTPError
	var newTodo dto.TodosRequestBody

	newTodo, err = getDTOCreateRequest(c, []string{"nama", "tanggal_mulai", "tanggal_selesai", "deskripsi"})
	if err != nil {
		httpError = dto.NewHTTPError(httpCode, "Invalid Request")
		goto ERROR_HANDLER
	}

	err = ctr.service.Create(newTodo)
	if err != nil {
		httpError = dto.NewHTTPError(httpCode, "Failed to create data")
		goto ERROR_HANDLER
	}

	c.AbortWithStatus(http.StatusOK)
	return
ERROR_HANDLER:
	if err != nil {
		c.JSON(httpError.ErrCode, httpError)
	}
	return
}

// ListAccounts godoc
// @Summary      Update one todos
// @Description  Update one todos based on the ID and JSON supplied by the user
// @Tags         todos
// @Accept       json
// @param id path int true "ID of todos to be updated"
// @Success      200  {object} dto.TodosResponseBody
// @Failure      400  {object}  dto.HTTPError
// @Failure      404  {object}  dto.HTTPError
// @Router       /todos/:id [put]
func (ctr *controller) Update(c *gin.Context) {
	var err error
	var httpCode int = http.StatusBadRequest
	var httpError dto.HTTPError
	var newTodo dto.TodosRequestBody
	var id string = c.Param("id")

	validId, err := validateId(id)

	newTodo, err = getDTOCreateRequest(c, []string{"nama", "tanggal_mulai", "tanggal_selesai", "deskripsi"})
	if err != nil {
		httpError = dto.NewHTTPError(httpCode, err.Error())
		goto ERROR_HANDLER
	}

	_, err = ctr.service.GetOne(validId)
	if err != nil {
		httpError = dto.NewHTTPError(http.StatusNotFound, err.Error())
		goto ERROR_HANDLER
	}
	err = ctr.service.Update(newTodo, validId)
	if err != nil {
		httpError = dto.NewHTTPError(httpCode, err.Error())
		goto ERROR_HANDLER
	}

	c.AbortWithStatus(http.StatusOK)
	return
ERROR_HANDLER:
	if err != nil {
		c.JSON(httpError.ErrCode, httpError)
	}
	return
}

// ListAccounts godoc
// @Summary      Delete one todos
// @Description  Delete one todos base don the ID supplied by the user
// @Tags         todos
// @param id path int true "ID of todos to be deleted"
// @Success      200  {object} dto.TodosResponseBody
// @Failure      404  {object}  dto.HTTPError
// @Failure      400  {object}  dto.HTTPError
// @Router       /todos/:id [delete]
func (ctr *controller) Delete(c *gin.Context) {
	var id string = c.Param("id")
	var err error
	var httpError dto.HTTPError
	validId, err := validateId(id)

	_, err = ctr.service.GetOne(validId)
	if err != nil {
		httpError = dto.NewHTTPError(http.StatusNotFound, err.Error())
		goto ERROR_HANDLER
	}

	err = ctr.service.Delete(validId)
	if err != nil {
		httpError = dto.NewHTTPError(http.StatusBadRequest, err.Error())
		goto ERROR_HANDLER
	}

	c.AbortWithStatus(http.StatusOK)
	return
ERROR_HANDLER:
	if err != nil {
		c.JSON(httpError.ErrCode, httpError)
	}
	return
}
