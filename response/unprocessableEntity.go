package response

func UnprocessableEntity(message string) (status int16, res commonResponse) {
	meg := "Unprocessable Entity"
	if message != "" {
		meg = message
	}
	res = commonResponse{
		meg,
	}
	return 422, res
}
