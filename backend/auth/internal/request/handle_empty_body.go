package request

func HandleEmptyBodyFields(fields ...string) bool {
	for _, field := range fields {
		if field == "" {
			return true
		}
	}

	return false
}
