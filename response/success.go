package response

func Success(item interface{}, message string) (status int, res ItemResponse) {
	meg := "Success"
	if message != "" {
		meg = message
	}
	res = ItemResponse{
		meg,
		item,
	}
	return 200, res
}
