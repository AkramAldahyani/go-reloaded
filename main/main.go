package main

import (
	"fmt"
	"os"
	"car"
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
			
			fmt.Println(car.ArrayToString(data))

		}

	} else {
		fmt.Print("Incorrect argument format! Please try again.")
	}
}

// for i := 0; i < len(data); i++ {
// 	fmt.Println(data[i])
// }
