package utils

func Contains(slice []string, input string) bool {
	for _, s := range slice {
		if len(input) >= len(s) && input[:len(s)] == s {
			return true
		}
	}
	return false
}
