package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseUnauthorizedTestSuite struct {
	suite.Suite
	resExpected commonResponse
}

func (suite *ResponseUnauthorizedTestSuite) TestDefaultMessage() {
	suite.resExpected.Message = "Unauthorized"
	status, response := Unauthorized("")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 401, status)
}

func (suite *ResponseUnauthorizedTestSuite) TestCustomMessage() {
	message := "This user don't have permission for this level"
	suite.resExpected.Message = message
	status, response := Unauthorized(message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 401, status)
}

func TestResponseUnauthorizedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseUnauthorizedTestSuite))
}
