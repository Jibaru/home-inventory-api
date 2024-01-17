package responses

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDataResponse(t *testing.T) {
	expectedResponse := DataResponse{
		Data: "test",
	}

	actualResponse := NewDataResponse("test")
	assert.Equal(t, expectedResponse.Data, actualResponse.Data)

	expectedJSON, err := json.Marshal(expectedResponse)
	assert.NoError(t, err)

	actualJSON, err := json.Marshal(actualResponse)
	assert.NoError(t, err)
	assert.JSONEq(t, string(expectedJSON), string(actualJSON))
}
