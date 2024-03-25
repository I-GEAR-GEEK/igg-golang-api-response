package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseInternalServerErrorTestSuite struct {
	suite.Suite
}

func (suite *ResponseInternalServerErrorTestSuite) TestDefaultMessage() {
	status, response := InternalServerError("", "", "")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Internal server error","code":""}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 500, status)
}

func (suite *ResponseInternalServerErrorTestSuite) TestCustomMessage() {
	status, response := InternalServerError("Your request server error", "", "GET_PROFILE_ERROR")
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Your request server error","code":"GET_PROFILE_ERROR"}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 500, status)
}

func TestResponseInternalServerErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseInternalServerErrorTestSuite))
}
