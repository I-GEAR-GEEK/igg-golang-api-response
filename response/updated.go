package response

func Updated(item interface{}, message string) (status int, res ItemResponse) {
	meg := "Updated successfully"
	if message != "" {
		meg = message
	}
	res = ItemResponse{
		meg,
		item,
	}
	return 202, res
}
