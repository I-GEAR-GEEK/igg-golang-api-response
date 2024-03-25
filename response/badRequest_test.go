package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseBadRequestTestSuite struct {
	suite.Suite
}

func (suite *ResponseBadRequestTestSuite) TestDefaultMessage() {
	status, response := BadRequest("", "")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Bad request","code":""}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 400, status)
}

func (suite *ResponseBadRequestTestSuite) TestCustomMessage() {
	status, response := BadRequest("Your request is bad", "EMAIL_DUPLICATE")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Your request is bad","code":"EMAIL_DUPLICATE"}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 400, status)
}

func TestResponseBadRequestTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseBadRequestTestSuite))
}
