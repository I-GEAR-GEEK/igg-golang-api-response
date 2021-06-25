package response

func BadRequest(message string) (status int, res commonResponse) {
	meg := "Bad request"
	if message != "" {
		meg = message
	}
	res = commonResponse{
		meg,
	}
	return 400, res
}
