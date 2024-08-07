//this function converts the modified array to string again.

package car

func ArrayToString(arr []string) string {
	var str string
	for i := 0; i < len(arr); i++ {
		r := []rune(arr[i])
		if r[0] == '(' {

		}else if arr[i] == "," || arr[i] == "."  || arr[i] == ":" || arr[i] == "!" || arr[i] == "?" {
			str += arr[i]
		}else if i == 0 {
			str += arr[i]
		}else{
			str += " " + arr[i]
		}
	}
	return str
}