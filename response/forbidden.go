package response

func Forbidden(message string, code string) (status int, res ErrorResponse) {
	if message == "" {
		message = "Forbidden"
	}
	res = ErrorResponse{
		Message: message,
		Code:    code,
	}
	return 403, res
}
