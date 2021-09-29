package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseNotFoundTestSuite struct {
	suite.Suite
	resExpected CommonResponse
}

func (suite *ResponseNotFoundTestSuite) TestDefaultMessage() {
	suite.resExpected.Message = "Not Found"
	status, response := NotFound("")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 404, status)
}

func (suite *ResponseNotFoundTestSuite) TestCustomMessage() {
	message := "Your request not found"
	suite.resExpected.Message = message
	status, response := NotFound(message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 404, status)
}

func TestResponseNotFoundTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseNotFoundTestSuite))
}
