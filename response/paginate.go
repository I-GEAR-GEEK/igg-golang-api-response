package response

import "math"

func Paginate(message string, data Pagination) (status int, res PaginationResponse) {
	if message == "" {
		message = "Data retrieval successfully"
	}
	d := float64(data.Total) / float64(data.PerPage)
	totalPage := int64(math.Ceil(d))
	res = PaginationResponse{
		Message:   message,
		Data:      data.Data,
		Meta:      data.Meta,
		Page:      data.Page,
		PerPage:   data.PerPage,
		Total:     data.Total,
		TotalPage: totalPage,
		Next:      data.Page + 1,
		Prev:      data.Page - 1,
	}
	return 200, res
}
