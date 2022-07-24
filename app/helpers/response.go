package helpers

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyResponse struct{}

func BuildResponse(status bool, message string, error interface{}, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   error,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitErr := strings.Split(err, "\n")

	res := Response{
		Status:  false,
		Message: message,
		Error:   splitErr,
		Data:    data,
	}

	return res
}
