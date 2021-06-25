package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type ResponseItemTestSuite struct {
	suite.Suite
	item        interface{}
	resExpected itemResponse
}

func (suite *ResponseItemTestSuite) SetupTest() {
	suite.item = bson.M{"productName": "table"}
}

func (suite *ResponseItemTestSuite) TestDefaultMessage() {
	suite.resExpected.Data = suite.item
	suite.resExpected.Message = "Data retrieval successfully"
	status, response := Item(suite.item, "")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 200, status)
}

func (suite *ResponseItemTestSuite) TestCustomMessage() {
	message := "Your request is complete"
	suite.resExpected.Data = suite.item
	suite.resExpected.Message = message
	status, response := Item(suite.item, message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 200, status)
}

func TestResponseItemTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseItemTestSuite))
}
