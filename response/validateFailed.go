package response

func ValidateFailed(errorMessages interface{}, message string) (status int, res ValidateFail) {
	meg := "Validation failed"
	if message != "" {
		meg = message
	}
	res = ValidateFail{
		meg,
		errorMessages,
	}
	return 422, res
}
