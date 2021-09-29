package response

func NotFound(message string) (status int, res CommonResponse) {
	meg := "Not Found"
	if message != "" {
		meg = message
	}
	res = CommonResponse{
		meg,
	}
	return 404, res
}
