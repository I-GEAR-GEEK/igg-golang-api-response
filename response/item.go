package response

func Item(item interface{}, message string) (status int, res itemResponse) {
	meg := "Data retrieval successfully"
	if message != "" {
		meg = message
	}
	res = itemResponse{
		meg,
		item,
	}
	return 200, res
}
