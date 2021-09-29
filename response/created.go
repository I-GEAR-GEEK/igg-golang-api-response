package response

func Created(item interface{}, message string) (status int, res ItemResponse) {
	meg := "Created successfully"
	if message != "" {
		meg = message
	}
	res = ItemResponse{
		meg,
		item,
	}
	return 201, res
}
