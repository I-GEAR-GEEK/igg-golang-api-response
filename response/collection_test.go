package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type ResponseCollectionTestSuite struct {
	suite.Suite
	defaultMessage string
	items          []interface{}
	resExpected    collectionResponse
}

func (suite *ResponseCollectionTestSuite) SetupTest() {
	suite.defaultMessage = "Data retrieval successfully"
	suite.items = append(suite.items, bson.M{"productName": "table"})
}

func (suite *ResponseCollectionTestSuite) TestDefaultMessage() {
	suite.resExpected.Data = suite.items
	suite.resExpected.Message = "Data retrieval successfully"
	status, response := Collection(suite.items, "")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), int16(200), status)
}

func (suite *ResponseCollectionTestSuite) TestCustomMessage() {
	message := "Your request is complete"
	suite.resExpected.Data = suite.items
	suite.resExpected.Message = message
	status, response := Collection(suite.items, message)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), int16(200), status)
}

func TestResponseCollectionTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseCollectionTestSuite))
}
