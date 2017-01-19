package helpers

// IsAlpha returns true when given string is alphanumeric
func IsAlpha(c rune) bool {
	if c < 'A' || c > 'z' {
		return false
	} else if c > 'Z' && c < 'a' {
		return false
	}

	return true
}
