package car

func StringToArray(s string) []string {
	var word string
	var arr []string
	data := string(s) + " "
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "(" {
			exit := false
			for !exit && i < len(data) {
				word += string(data[i])
				if string(data[i]) == ")" {
					exit = true
				}
				i++
			}
			if word != "" {
				arr = append(arr, word)
				word = ""
			}

		} else if string(data[i]) == "," || string(data[i]) == "." || string(data[i]) == "!" || string(data[i]) == "?" || string(data[i]) == ":" || string(data[i]) == "'" {
			if word != "" {
				arr = append(arr, word)
				word = ""
			}
			word = string(data[i])
			arr = append(arr, word)
			word = ""
		} else if string(data[i]) != " " {
			word += string(data[i])
		} else {
			if word != "" {
				arr = append(arr, word)
				word = ""
			}
			word = ""
		}
	}
	return arr
}
