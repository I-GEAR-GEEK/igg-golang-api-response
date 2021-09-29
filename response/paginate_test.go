package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponsePaginateTestSuite struct {
	suite.Suite
	resExpected PaginationResponse
}

func (suite *ResponsePaginateTestSuite) SetupTest() {
	suite.resExpected.Page = 1
	suite.resExpected.PerPage = 10
	suite.resExpected.Total = 100
	suite.resExpected.TotalPage = 10
	suite.resExpected.Next = 2
	suite.resExpected.Prev = 0
}

func (suite *ResponsePaginateTestSuite) TestDefaultMessageAnd100Row10Perpage() {
	items := []interface{}{"data 1", "data 2"}
	data := Pagination{items, 100, 10, 1}
	suite.resExpected.Message = "Data retrieval successfully"
	suite.resExpected.Data = items
	status, response := Paginate(data, "")
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 200, status)
}

func (suite *ResponsePaginateTestSuite) TestCustomMessageAnd1Row10Perpage() {
	customMessage := "Get paginate success"
	items := []interface{}{"data 1"}
	data := Pagination{items, 1, 10, 1}
	suite.resExpected.Message = customMessage
	suite.resExpected.Data = items
	suite.resExpected.Total = 1
	suite.resExpected.TotalPage = 1
	status, response := Paginate(data, customMessage)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 200, status)
}

func (suite *ResponsePaginateTestSuite) TestDefaultMessageAnd13Row10Perpage() {
	customMessage := "Get paginate success"
	items := []interface{}{"data 1"}
	data := Pagination{items, 13, 10, 1}
	suite.resExpected.Message = customMessage
	suite.resExpected.Data = items
	suite.resExpected.Total = 13
	suite.resExpected.TotalPage = 2
	status, response := Paginate(data, customMessage)
	assert.Equal(suite.T(), suite.resExpected, response)
	assert.Equal(suite.T(), 200, status)
}

func TestResponsePaginateTestSuite(t *testing.T) {
	suite.Run(t, new(ResponsePaginateTestSuite))
}
