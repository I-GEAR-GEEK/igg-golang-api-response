package response

func BadRequest(message string) (status int, res CommonResponse) {
	meg := "Bad request"
	if message != "" {
		meg = message
	}
	res = CommonResponse{
		meg,
	}
	return 400, res
}
