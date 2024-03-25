package response

func Deleted(message string) (status int, res CommonResponse) {
	if message == "" {
		message = "No content"
	}
	res = CommonResponse{
		Message: message,
	}
	return 204, res
}
