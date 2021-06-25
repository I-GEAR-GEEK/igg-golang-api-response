package response

func Unauthorized(message string) (status int, res commonResponse) {
	meg := "Unauthorized"
	if message != "" {
		meg = message
	}
	res = commonResponse{
		meg,
	}
	return 401, res
}
