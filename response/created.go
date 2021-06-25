package response

func Created(item interface{}, message string) (status int, res itemResponse) {
	meg := "Created successfully"
	if message != "" {
		meg = message
	}
	res = itemResponse{
		meg,
		item,
	}
	return 201, res
}
