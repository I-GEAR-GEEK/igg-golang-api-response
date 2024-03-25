package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseNotFoundTestSuite struct {
	suite.Suite
}

func (suite *ResponseNotFoundTestSuite) TestDefaultMessage() {
	status, response := NotFound("", "")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Not Found","code":""}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 404, status)
}

func (suite *ResponseNotFoundTestSuite) TestCustomMessage() {
	status, response := NotFound("Your request not found", "USER_NOT_FOUND")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Your request not found","code":"USER_NOT_FOUND"}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 404, status)
}

func TestResponseNotFoundTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseNotFoundTestSuite))
}
