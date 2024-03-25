package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseItemTestSuite struct {
	suite.Suite
}

func (suite *ResponseItemTestSuite) TestDefaultMessage() {
	status, response := Item("", "")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Data retrieval successfully","data":""}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 200, status)
}

func (suite *ResponseItemTestSuite) TestCustomMessage() {
	item := map[string]interface{}{
		"id": 1,
	}
	status, response := Item("Your request is complete", item)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Your request is complete","data":{"id":1}}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 200, status)
}

func TestResponseItemTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseItemTestSuite))
}
