package response

import "math"

func Paginate(data Pagination, message string) (status int, res PaginationResponse) {
	msg := "Data retrieval successfully"
	if message != "" {
		msg = message
	}
	d := float64(data.Total) / float64(data.PerPage)
	totalPage := int64(math.Ceil(d))
	res = PaginationResponse{
		msg,
		data.Data,
		data.Meta,
		data.Page,
		data.PerPage,
		data.Total,
		totalPage,
		data.Page + 1,
		data.Page - 1,
	}
	return 200, res
}
