package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseDeletedTestSuite struct {
	suite.Suite
	resExpected commonResponse
}

func (suite *ResponseDeletedTestSuite) TestDefaultMessage() {
	suite.resExpected.Message = "No content"
	status, response := Deleted("")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 204, status)
}

func (suite *ResponseDeletedTestSuite) TestCustomMessage() {
	message := "Your request is complete"
	suite.resExpected.Message = message
	status, response := Deleted(message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 204, status)
}

func TestResponseDeletedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseDeletedTestSuite))
}
