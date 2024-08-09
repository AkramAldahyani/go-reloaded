package reloaded
//This function is for the capitalaization.
func Capitalize(s string) string {
	str := []rune(s)
	start := true
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = str[i] + 32
		}
	}
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' || (str[i] >= 'A' && str[i] <= 'Z') {
			if start && (str[i] >= 'a' && str[i] <= 'z') {
				str[i] = str[i] - 32
				start = false
			}
		} else if str[i] >= '0' && str[i] <= '9' {
			start = false
		} else {
			start = true
		}
	}

	return string(str)
}
