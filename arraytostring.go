// this function converts the modified array to more orgnized string.

package reloaded

func ArrayToString(arr []string) string {
	var str string
	for i := 0; i < len(arr); i++ {
		if arr[i] == "" {
		} else if arr[i] == "," || arr[i] == "." || arr[i] == ":" || arr[i] == "!" || arr[i] == "?" {
			str += arr[i]
		} else if arr[i] == "'" {
			exit := false
			if i != 0 {
				str += " "
			}
			str += "'"
			j := 0
			i++
			for !exit && i < len(arr)-1 {

				if j == 0 {
					str += arr[i]
				} else if arr[i] == "," || arr[i] == "." || arr[i] == ":" || arr[i] == "!" || arr[i] == "?" {
					str += arr[i]
				} else {
					str += " " + arr[i]
				}
				i++
				j++
				if arr[i] == "'" {
					exit = true
					str += "'"
				}
			}
		} else if i == 0{
			str += arr[i]
		}else {
			str += " " + arr[i]
		}
	}
	return str
}
