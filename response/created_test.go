package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseCreatedTestSuite struct {
	suite.Suite
}

func (suite *ResponseCreatedTestSuite) TestDefaultMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := Created("Created successfully", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Created successfully","data":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 201, status)
}

func (suite *ResponseCreatedTestSuite) TestCustomMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := Created("Your request is complete", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Your request is complete","data":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 201, status)
}

func TestResponseCreatedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseCreatedTestSuite))
}
