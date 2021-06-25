package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseInternalServerErrorTestSuite struct {
	suite.Suite
	resExpected commonResponse
}

func (suite *ResponseInternalServerErrorTestSuite) TestDefaultMessage() {
	suite.resExpected.Message = "Internal server error"
	status, response := InternalServerError("")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), int16(500), status)
}

func (suite *ResponseInternalServerErrorTestSuite) TestCustomMessage() {
	message := "Your request server error"
	suite.resExpected.Message = message
	status, response := InternalServerError(message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), int16(500), status)
}

func TestResponseInternalServerErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseInternalServerErrorTestSuite))
}
