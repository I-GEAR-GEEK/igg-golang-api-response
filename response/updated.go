package response

func Updated(item interface{}, message string) (status int, res itemResponse) {
	meg := "Updated successfully"
	if message != "" {
		meg = message
	}
	res = itemResponse{
		meg,
		item,
	}
	return 202, res
}
