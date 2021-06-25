package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type ResponseCreatedTestSuite struct {
	suite.Suite
	item        interface{}
	resExpected itemResponse
}

func (suite *ResponseCreatedTestSuite) SetupTest() {
	suite.item = bson.M{"productName": "table"}
}

func (suite *ResponseCreatedTestSuite) TestDefaultMessage() {
	suite.resExpected.Data = suite.item
	suite.resExpected.Message = "Created successfully"
	status, response := Created(suite.item, "")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 201, status)
}

func (suite *ResponseCreatedTestSuite) TestCustomMessage() {
	message := "Your request is complete"
	suite.resExpected.Data = suite.item
	suite.resExpected.Message = message
	status, response := Created(suite.item, message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 201, status)
}

func TestResponseCreatedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseCreatedTestSuite))
}
