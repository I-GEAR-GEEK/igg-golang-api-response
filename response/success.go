package response

func Success(item interface{}, message string) (status int16, res itemResponse) {
	meg := "Success"
	if message != "" {
		meg = message
	}
	res = itemResponse{
		meg,
		item,
	}
	return 200, res
}
