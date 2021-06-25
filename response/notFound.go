package response

func NotFound(message string) (status int, res commonResponse) {
	meg := "Not Found"
	if message != "" {
		meg = message
	}
	res = commonResponse{
		meg,
	}
	return 404, res
}
