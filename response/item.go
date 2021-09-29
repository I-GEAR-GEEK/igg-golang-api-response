package response

func Item(item interface{}, message string) (status int, res ItemResponse) {
	meg := "Data retrieval successfully"
	if message != "" {
		meg = message
	}
	res = ItemResponse{
		meg,
		item,
	}
	return 200, res
}
