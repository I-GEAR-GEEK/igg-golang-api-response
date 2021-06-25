package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type ResponseValidateFailedTestSuite struct {
	suite.Suite
	errorMessages []interface{}
	resExpected   validateFailed
}

func (suite *ResponseValidateFailedTestSuite) SetupTest() {
	suite.errorMessages = append(suite.errorMessages, bson.M{"fieldName": "productName"})
}

func (suite *ResponseValidateFailedTestSuite) TestDefaultMessage() {
	suite.resExpected.Errors = suite.errorMessages
	suite.resExpected.Message = "Validation failed"
	status, response := ValidateFailed(suite.errorMessages, "")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 422, status)
}

func (suite *ResponseValidateFailedTestSuite) TestCustomMessage() {
	message := "Some field error"
	suite.resExpected.Errors = suite.errorMessages
	suite.resExpected.Message = message
	status, response := ValidateFailed(suite.errorMessages, message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 422, status)
}

func TestResponseValidateFailedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseValidateFailedTestSuite))
}
