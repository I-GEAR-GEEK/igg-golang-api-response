package response

func ValidateFailed(errorMessages []interface{}, message string) (status int16, res validateFailed) {
	meg := "Validation failed"
	if message != "" {
		meg = message
	}
	res = validateFailed{
		meg,
		errorMessages,
	}
	return 422, res
}