package main

import (
	"car"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 2 {
		dataFile := files[0]
		dat, err := os.ReadFile(dataFile)
		if err != nil {
			fmt.Println("Please enter a valid file.")
			return
		} else {
			data := car.StringToArray(string(dat))
			for i := 0; i < len(data); i++ {
				switch data[i] {
				case "(cap)":
					data[i-1] = car.Capitalize(data[i-1])
				case "(up)":
					data[i-1] = car.ToUpper(data[i-1])
				case "(low)":
					data[i-1] = car.ToLower(data[i-1])
				}
			}
			fmt.Println(car.ArrayToString(data))
			
		}

	} else {
		fmt.Print("Incorrect argument format! Please try again.")
	}

}

// for i := 0; i < len(data); i++ {
// 	fmt.Println(data[i])
// }
