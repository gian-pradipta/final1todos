package controller

import (
	"encoding/json"
	"final_satu/internal/dto"
	"final_satu/internal/helper"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

func validateId(id string) (int, error) {
	var err error
	var validatedId int

	validatedId, err = strconv.Atoi(id)

	return validatedId, err
}

func getDTOCreateRequest(c *gin.Context, allowedParams []string) (dto.TodosRequestBody, error) {
	var err error
	var createRequest dto.TodosRequestBody
	jsonByte, err := io.ReadAll(c.Request.Body)
	// jsonByte, err := json.Marshal(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		return createRequest, err
	}

	err = helper.ValidateJSONParams(jsonByte, allowedParams)
	if err != nil {
		return createRequest, err
	}

	err = json.Unmarshal(jsonByte, &createRequest)
	if err != nil {
		return createRequest, err
	}

	err = helper.ValidateJsonCreateRequest(&createRequest)
	if err != nil {
		return createRequest, err
	}

	return createRequest, err

}
