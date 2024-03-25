package response

func Created(message string, item interface{}) (status int, res ItemResponse) {
	if message == "" {
		message = "Created successfully"
	}
	res = ItemResponse{
		Message: message,
		Data:    item,
	}
	return 201, res
}
