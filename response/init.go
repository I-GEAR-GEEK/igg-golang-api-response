package response

type CommonResponse struct {
	Message string `json:"message" bson:"message"`
}

type ValidateFail struct {
	Message string      `json:"message" bson:"message"`
	Errors  interface{} `json:"errors" bson:"errors"`
}

type ItemResponse struct {
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type CollectionResponse struct {
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type PaginationResponse struct {
	Message   string      `json:"message" bson:"message"`
	Data      interface{} `json:"data" bson:"data"`
	Meta      interface{} `json:"meta" bson:"meta"`
	Page      int64       `json:"page" bson:"page"`
	PerPage   int64       `json:"per_page" bson:"per_page"`
	Total     int64       `json:"total" bson:"total"`
	TotalPage int64       `json:"total_page" bson:"total_page"`
	Next      int64       `json:"next" bson:"next"`
	Prev      int64       `json:"prev" bson:"prev"`
}

type Pagination struct {
	Data    interface{} `json:"data" bson:"data"`
	Meta    interface{} `json:"meta" bson:"meta"`
	Total   int64       `json:"total" bson:"total"`
	PerPage int64       `json:"per_page" bson:"per_page"`
	Page    int64       `json:"page" bson:"page"`
}
