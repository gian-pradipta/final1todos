package helper

import (
	"encoding/json"
	"errors"
	"final_satu/internal/dto"
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateJSONParams(jsonByte []byte, allowedParams []string) error {
	var err error

	var validationMap map[string]interface{} = make(map[string]interface{})
	err = json.Unmarshal(jsonByte, &validationMap)
	if err != nil {
		return err
	}

	var allowedParamsMap map[string]bool = make(map[string]bool)
	for _, param := range allowedParams {
		allowedParamsMap[param] = true
	}

	for key := range validationMap {
		if allowedParamsMap[key] == false {
			err = errors.New("Invalid JSON Parameters")
			return err
		}
	}

	if len(allowedParams) != len(validationMap) {
		err = errors.New("Invalid JSON Parameters")
		return err
	}

	return err
}

func ValidateJsonCreateRequest(jsonStruct *dto.TodosRequestBody) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(jsonStruct)
	layout := "2006-01-02"
	dateString := jsonStruct.TanggalMulai
	dateString2 := jsonStruct.TanggalSelesai
	d1, err := time.Parse(layout, dateString)
	if err != nil {
		return err
	}

	d2, err := time.Parse(layout, dateString2)
	if err != nil {
		return err
	}

	if !(d1.Equal(d2) || d1.Before(d2)) {
		return errors.New("Tanggal Selesai Invalid")
	}
	return err
}
