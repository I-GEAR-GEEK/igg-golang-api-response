package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseAcceptedTestSuite struct {
	suite.Suite
}

func (suite *ResponseAcceptedTestSuite) TestDefaultMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := Accepted("Accepted successfully", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Accepted successfully","data":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 202, status)
}

func (suite *ResponseAcceptedTestSuite) TestCustomMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := Accepted("Your request is complete", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Your request is complete","data":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 202, status)
}

func TestResponseAcceptedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseAcceptedTestSuite))
}
