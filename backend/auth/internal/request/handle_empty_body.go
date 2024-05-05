package request

func HandleEmptyBody(fields ...string) bool {
	for _, field := range fields {
		if field == "" {
			return true
		}
	}

	return false
}
