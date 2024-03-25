package response

func NotFound(message string, code string) (status int, res ErrorResponse) {
	if message == "" {
		message = "Not Found"
	}
	res = ErrorResponse{
		Message: message,
		Code:    code,
	}
	return 404, res
}
