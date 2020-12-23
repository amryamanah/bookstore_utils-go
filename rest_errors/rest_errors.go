package rest_errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message    string        `json:"message"`
	StatusCode int           `json:"status_code"`
	Error      string        `json:"error"`
	Causes     []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(m string) *RestErr {
	return &RestErr{
		Message:    m,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

func NewNotFoundError(m string) *RestErr {
	return &RestErr{
		Message:    m,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

func NewInternalServerError(m string, err error) *RestErr {
	result := &RestErr{
		Message:    m,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
