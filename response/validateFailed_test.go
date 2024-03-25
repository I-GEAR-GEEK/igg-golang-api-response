package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseValidateFailedTestSuite struct {
	suite.Suite
}

func (suite *ResponseValidateFailedTestSuite) TestDefaultMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := ValidateFailed("Validation failed", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Validation failed","errors":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 422, status)
}

func (suite *ResponseValidateFailedTestSuite) TestCustomMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := ValidateFailed("Some field error", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Some field error","errors":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 422, status)
}

func TestResponseValidateFailedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseValidateFailedTestSuite))
}
