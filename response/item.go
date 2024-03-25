package response

func Item(message string, item interface{}) (status int, res ItemResponse) {
	if message == "" {
		message = "Data retrieval successfully"
	}
	res = ItemResponse{
		Message: message,
		Data:    item,
	}
	return 200, res
}
