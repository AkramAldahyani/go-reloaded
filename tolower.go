package reloaded

// Turn the string to all lower case strings
func ToLower(s string) string {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}

// Turn the string to all upper case letters.
func ToUpper(s string) string {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] = str[i] - 32
		}
	}
	return string(str)
}
