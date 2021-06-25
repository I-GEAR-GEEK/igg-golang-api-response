package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type ResponseSuccessTestSuite struct {
	suite.Suite
	item        interface{}
	resExpected itemResponse
}

func (suite *ResponseSuccessTestSuite) SetupTest() {
	suite.item = bson.M{"productName": "table"}
}

func (suite *ResponseSuccessTestSuite) TestDefaultMessage() {
	suite.resExpected.Data = suite.item
	suite.resExpected.Message = "Success"
	status, response := Success(suite.item, "")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 200, status)
}

func (suite *ResponseSuccessTestSuite) TestCustomMessage() {
	message := "Your request is success"
	suite.resExpected.Data = suite.item
	suite.resExpected.Message = message
	status, response := Success(suite.item, message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 200, status)
}

func TestResponseSuccessTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseSuccessTestSuite))
}
