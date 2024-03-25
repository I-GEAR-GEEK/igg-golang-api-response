package response

func Updated(message string, item interface{}) (status int, res ItemResponse) {
	if message == "" {
		message = "Updated successfully"
	}
	res = ItemResponse{
		Message: message,
		Data:    item,
	}
	return 200, res
}
