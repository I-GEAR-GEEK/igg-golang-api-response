package response

func Success(message string, item interface{}) (status int, res ItemResponse) {
	if message == "" {
		message = "Success"
	}
	res = ItemResponse{
		Message: message,
		Data:    item,
	}
	return 200, res
}
