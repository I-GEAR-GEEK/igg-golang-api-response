package response

func Unauthorized(message string) (status int, res CommonResponse) {
	meg := "Unauthorized"
	if message != "" {
		meg = message
	}
	res = CommonResponse{
		meg,
	}
	return 401, res
}
