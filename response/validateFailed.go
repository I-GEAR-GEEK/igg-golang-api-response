package response

func ValidateFailed(message string, errorMessages interface{}) (status int, res ValidateFail) {
	if message == "" {
		message = "Validation failed"
	}
	res = ValidateFail{
		Message: message,
		Errors:  errorMessages,
	}
	return 422, res
}
