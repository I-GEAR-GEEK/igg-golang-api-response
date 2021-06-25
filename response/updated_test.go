package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type ResponseUpdatedTestSuite struct {
	suite.Suite
	item        interface{}
	resExpected itemResponse
}

func (suite *ResponseUpdatedTestSuite) SetupTest() {
	suite.item = bson.M{"productName": "table"}
}

func (suite *ResponseUpdatedTestSuite) TestDefaultMessage() {
	suite.resExpected.Data = suite.item
	suite.resExpected.Message = "Updated successfully"
	status, response := Updated(suite.item, "")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 202, status)
}

func (suite *ResponseUpdatedTestSuite) TestCustomMessage() {
	message := "Your request is success"
	suite.resExpected.Data = suite.item
	suite.resExpected.Message = message
	status, response := Updated(suite.item, message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 202, status)
}

func TestResponseUpdatedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseUpdatedTestSuite))
}
