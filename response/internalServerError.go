package response

func InternalServerError(message string, detail string, code string) (status int, res ErrorResponse) {
	if message == "" {
		message = "Internal server error"
	}
	res = ErrorResponse{
		Message: message,
		Detail:  detail,
		Code:    code,
	}
	return 500, res
}
