package response

func InternalServerError(message string) (status int, res CommonResponse) {
	meg := "Internal server error"
	if message != "" {
		meg = message
	}
	res = CommonResponse{
		meg,
	}
	return 500, res
}
