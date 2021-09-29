package response

func Deleted(message string) (status int, res CommonResponse) {
	meg := "No content"
	if message != "" {
		meg = message
	}
	res = CommonResponse{
		meg,
	}
	return 204, res
}
