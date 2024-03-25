package response

import (
	"encoding/json"
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
	meta := map[string]interface{}{
		"price_total": 10,
	}
	data := Pagination{items, meta, 100, 10, 1}
	status, response := Paginate("", data)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Data retrieval successfully","data":["data 1","data 2"],"meta":{"price_total":10},"page":1,"per_page":10,"total":100,"total_page":10,"next":2,"prev":0}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 200, status)
}

func (suite *ResponsePaginateTestSuite) TestCustomMessageAnd1Row10Perpage() {
	customMessage := "Get paginate success"
	items := []interface{}{"data 1"}
	meta := map[string]interface{}{
		"price_total": 10,
	}
	data := Pagination{items, meta, 1, 10, 1}
	status, response := Paginate(customMessage, data)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Get paginate success","data":["data 1"],"meta":{"price_total":10},"page":1,"per_page":10,"total":1,"total_page":1,"next":2,"prev":0}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 200, status)
}

func (suite *ResponsePaginateTestSuite) TestDefaultMessageAnd13Row10Perpage() {
	customMessage := "Get paginate success"
	items := []interface{}{"data 1"}
	meta := map[string]interface{}{
		"price_total": 10,
	}
	data := Pagination{items, meta, 13, 10, 1}
	status, response := Paginate(customMessage, data)
	resJson, _ := json.Marshal(response)
	expected := `{"message":"Get paginate success","data":["data 1"],"meta":{"price_total":10},"page":1,"per_page":10,"total":13,"total_page":2,"next":2,"prev":0}`
	assert.Equal(suite.T(), expected, string(resJson))
	assert.Equal(suite.T(), 200, status)
}

func TestResponsePaginateTestSuite(t *testing.T) {
	suite.Run(t, new(ResponsePaginateTestSuite))
}
