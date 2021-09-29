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
	Page      int64       `json:"page" bson:"page"`
	PerPage   int64       `json:"perPage" bson:"perPage"`
	Total     int64       `json:"total" bson:"total"`
	TotalPage int64       `json:"totalPage" bson:"totalPage"`
	Next      int64       `json:"next" bson:"next"`
	Prev      int64       `json:"prev" bson:"prev"`
}

type Pagination struct {
	Data    interface{} `json:"data" bson:"data"`
	Total   int64       `json:"total" bson:"total"`
	PerPage int64       `json:"perPage" bson:"perPage"`
	Page    int64       `json:"page" bson:"page"`
}
