package response

func InternalServerError(message string) (status int, res commonResponse) {
	meg := "Internal server error"
	if message != "" {
		meg = message
	}
	res = commonResponse{
		meg,
	}
	return 500, res
}
