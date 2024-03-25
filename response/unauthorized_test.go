package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseUnauthorizedTestSuite struct {
	suite.Suite
}

func (suite *ResponseUnauthorizedTestSuite) TestDefaultMessage() {
	status, response := Unauthorized("", "")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Unauthorized","code":""}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 401, status)
}

func (suite *ResponseUnauthorizedTestSuite) TestCustomMessage() {
	status, response := Unauthorized("This user don't have permission for this level", "PERMISSION_DENIED")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"This user don't have permission for this level","code":"PERMISSION_DENIED"}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 401, status)
}

func TestResponseUnauthorizedTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseUnauthorizedTestSuite))
}
