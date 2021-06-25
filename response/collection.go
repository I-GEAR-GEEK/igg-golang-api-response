package response

func Collection(items []interface{}, message string) (status int16, res collectionResponse) {
	meg := "Data retrieval successfully"
	if message != "" {
		meg = message
	}
	res = collectionResponse{
		meg,
		items,
	}
	return 200, res
}
