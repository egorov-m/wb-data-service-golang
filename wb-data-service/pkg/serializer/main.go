package serializer

import (
	"encoding/json"
	"wb-data-service-golang/wb-data-service/pkg/byteconv"
)

type _Response struct {
	IsSuccess bool `json:"is_success"`
	Data      any  `json:"data"`
}

type _Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type _ErrorResponse struct {
	IsSuccess bool `json:"is_success"`
	Error     _Error
	Data      any `json:"data"`
}

func SuccessResponse(data any) string {
	response := _Response{
		IsSuccess: true,
		Data:      data,
	}
	return getJsonBody(response)
}

func ErrorResponse(statusCode int, err error, data any) string {
	response := _ErrorResponse{
		IsSuccess: false,
		Error: _Error{
			StatusCode: statusCode,
			Message:    err.Error(),
		},
		Data: data,
	}
	return getJsonBody(response)
}

func getJsonBody(response any) string {
	jsonResult, err := json.Marshal(response)
	if err != nil {
		return err.Error()
	}

	return byteconv.String(jsonResult)
}
