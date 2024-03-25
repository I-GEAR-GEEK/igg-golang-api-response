package response

func Unauthorized(message string, code string) (status int, res ErrorResponse) {
	if message == "" {
		message = "Unauthorized"
	}
	res = ErrorResponse{
		Message: message,
		Code:    code,
	}
	return 401, res
}
