/*if there is a single panctuation and no spaces before or after then it will be a single word, but if there any space so its not a single word*/
package reloaded
//This function is used to turn the string to array based on a several rules.
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

		} else if i > 0 && i < len(data)-1 && string(data[i]) == "'" && string(data[i-1]) != " " && string(data[i+1]) != " " {
			word += string(data[i])
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
