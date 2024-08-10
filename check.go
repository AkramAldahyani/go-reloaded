/*this func is used to change the format of the input where every condeition will be implemnted and deleted at the same time, so we can get a clear output*/
package reloaded

import (
	"regexp"
	"strconv"
)

func Check(arr []string) []string {
	//cheking in the array for the conditions.
	var input string
	wordRegex := regexp.MustCompile(`\b(bin|cap|hex|low|up)\b`)
	numberRegex := regexp.MustCompile(`\d+`)
	for i := 0; i < len(arr); i++ {
		input = arr[i]
		word := wordRegex.FindString(input)
		number := numberRegex.FindString(input)
		num, err := strconv.Atoi(number)
		if err != nil {
			num = 1
		}
		switch word {
		case "cap":
			arr = Reverse(Capitalize, arr, num, i-1)
			arr[i] = ""
		case "low":
			arr = Reverse(ToLower, arr, num, i-1)
			arr[i] = ""
		case "up":
			arr = Reverse(ToUpper, arr, num, i-1)
			arr[i] = ""
		case "bin":
			arr = Reverse(Bin, arr, num, i-1)
			arr[i] = ""
		case "hex":
			arr = Reverse(Hex, arr, num, i-1)
			arr[i] = ""
		}
	}

	return arr

}
func Reverse(f func(string) string, arr []string, num, pos int) []string {
	//this function for implementing the function over the arrays elements
	j := 1
	for i := pos; i >= 0; i-- {
		if arr[i] == "" {
			continue
		}
		if j <= num {
			arr[i] = f(arr[i])
			j++
		}
		if j > num {
			break
		}
	}
	return arr
}
func VowelLetter(arr []string) []string {
	//This function will check for the vowel letters.
	vowel := regexp.MustCompile(`^[AEIOUaeiou]`)

	for i := 1; i < len(arr); i++ {
		if vowel.MatchString(arr[i]) {
			if arr[i-1] == "a" {
				arr[i-1] = "an"
			} else if arr[i-1] == "A" {
				arr[i-1] = "An"
			}

		}

	}

	return arr
}

func Bin(b string) string {
	//binary to integer
	decimalValue, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return "Error"
	}
	value := strconv.FormatInt(decimalValue, 10)
	return string(value)
}

func Hex(h string) string {
	//hex to integer
	hexValue, err := strconv.ParseInt(h, 16, 64)
	if err != nil {
		return "Error"
	}
	value := strconv.FormatInt(hexValue, 10)
	return string(value)
}
