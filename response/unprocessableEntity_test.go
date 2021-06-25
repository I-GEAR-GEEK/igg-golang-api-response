package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseUnprocessableEntityTestSuite struct {
	suite.Suite
	resExpected commonResponse
}

func (suite *ResponseUnprocessableEntityTestSuite) TestDefaultMessage() {
	suite.resExpected.Message = "Unprocessable Entity"
	status, response := UnprocessableEntity("")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), int16(422), status)
}

func (suite *ResponseUnprocessableEntityTestSuite) TestCustomMessage() {
	message := "Something is wrong"
	suite.resExpected.Message = message
	status, response := UnprocessableEntity(message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), int16(422), status)
}

func TestResponseUnprocessableEntityTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseUnprocessableEntityTestSuite))
}
