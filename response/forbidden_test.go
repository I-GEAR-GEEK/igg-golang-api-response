package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseForbiddenTestSuite struct {
	suite.Suite
}

func (suite *ResponseForbiddenTestSuite) TestDefaultMessage() {
	status, response := Forbidden("", "")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Forbidden","code":""}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 403, status)
}

func (suite *ResponseForbiddenTestSuite) TestCustomMessage() {
	status, response := Forbidden("This user don't have permission for this level", "PERMISSION_DENIED")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"This user don't have permission for this level","code":"PERMISSION_DENIED"}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 403, status)
}

func TestResponseForbiddenTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseForbiddenTestSuite))
}
