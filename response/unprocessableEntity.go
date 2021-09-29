package response

func UnprocessableEntity(message string) (status int, res CommonResponse) {
	meg := "Unprocessable Entity"
	if message != "" {
		meg = message
	}
	res = CommonResponse{
		meg,
	}
	return 422, res
}
