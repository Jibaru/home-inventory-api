package responses

import (
	"encoding/json"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMessageResponse(t *testing.T) {
	message := random.String(100, random.Alphanumeric)
	expectedResponse := MessageResponse{
		Message: message,
	}

	actualResponse := NewMessageResponse(message)
	assert.Equal(t, expectedResponse.Message, actualResponse.Message)

	expectedJSON, err := json.Marshal(expectedResponse)
	assert.NoError(t, err)

	actualJSON, err := json.Marshal(actualResponse)
	assert.NoError(t, err)
	assert.JSONEq(t, string(expectedJSON), string(actualJSON))
}
