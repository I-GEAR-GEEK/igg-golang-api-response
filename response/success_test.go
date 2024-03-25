package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseSuccessTestSuite struct {
	suite.Suite
}

func (suite *ResponseSuccessTestSuite) TestDefaultMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := Success("Success", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Success","data":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 200, status)
}

func (suite *ResponseSuccessTestSuite) TestCustomMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := Success("Your request is success", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Your request is success","data":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 200, status)
}

func TestResponseSuccessTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseSuccessTestSuite))
}
