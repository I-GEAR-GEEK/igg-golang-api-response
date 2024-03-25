package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseUpdatedTestSuite struct {
	suite.Suite
}

func (suite *ResponseUpdatedTestSuite) TestDefaultMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := Updated("Updated successfully", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Updated successfully","data":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 200, status)
}

func (suite *ResponseUpdatedTestSuite) TestCustomMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := Updated("Your request is success", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Your request is success","data":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 200, status)
}

func TestResponseUpdatedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseUpdatedTestSuite))
}
