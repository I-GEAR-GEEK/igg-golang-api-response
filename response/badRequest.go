package response

func BadRequest(message string) (status int16, res commonResponse) {
	meg := "Bad request"
	if message != "" {
		meg = message
	}
	res = commonResponse{
		meg,
	}
	return 400, res
}
