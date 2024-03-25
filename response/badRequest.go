package response

func BadRequest(message string, code string) (status int, res ErrorResponse) {
	if message == "" {
		message = "Bad request"
	}
	res = ErrorResponse{
		Message: message,
		Code:    code,
	}
	return 400, res
}
