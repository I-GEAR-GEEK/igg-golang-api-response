package response

func Collection(items interface{}, message string) (status int, res CollectionResponse) {
	meg := "Data retrieval successfully"
	if message != "" {
		meg = message
	}
	res = CollectionResponse{
		meg,
		items,
	}
	return 200, res
}
