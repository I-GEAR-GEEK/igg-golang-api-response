package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseBadRequestTestSuite struct {
	suite.Suite
	resExpected commonResponse
}

func (suite *ResponseBadRequestTestSuite) SetupTest() {
	suite.resExpected.Message = "Bad request"
}

func (suite *ResponseBadRequestTestSuite) TestDefaultMessage() {
	status, response := BadRequest("")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), int16(400), status)
}

func (suite *ResponseBadRequestTestSuite) TestCustomMessage() {
	message := "Your request is bad"
	suite.resExpected.Message = message
	status, response := BadRequest(message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), int16(400), status)
}

func TestResponseBadRequestTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseBadRequestTestSuite))
}
