package rest_errors

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])

	errBytes, _ := json.Marshal(err)
	assert.EqualValues(t,
		"{\"message\":\"this is the message\",\"status_code\":500,\"error\":\"internal_server_error\",\"causes\":[\"database error\"]}",
		string(errBytes),
	)
}

func TestNewBadRequestError(t *testing.T) {

}
