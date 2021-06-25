package response

func Deleted(message string) (status int16, res commonResponse) {
	meg := "No content"
	if message != "" {
		meg = message
	}
	res = commonResponse{
		meg,
	}
	return 204, res
}
