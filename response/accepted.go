package response

func Accepted(message string, item interface{}) (status int, res ItemResponse) {
	if message == "" {
		message = "Accepted"
	}
	res = ItemResponse{
		Message: message,
		Data:    item,
	}
	return 202, res
}
