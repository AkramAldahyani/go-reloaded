/*this func will be used to change to format the input where every condeition will be implemnted and deleted at the same time*/
package car

import (
	"regexp"
	"strconv"
)

func Check(arr []string) []string {
	//var data []string
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
		}
	}

	return arr
}

func VowelLetter(arr []string) []string {
	//This function will check for the vowel letters.
	var data []string
	//compelete this function.
	return data
}

func Reverse(f func(string) string, arr []string, num, pos int) []string {
	//var data []string
	j := 0
	for i := pos; i >= 0 && num > 0; i-- {
		if i >= j {
			arr[i] = f(arr[i])
			j++
		}

		//fix the loop algorthim

	}
	return arr
}
