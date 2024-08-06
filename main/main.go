package main

import (
	"car"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 2{
		dataFile := files[0]
		da, err := os.ReadFile(dataFile)
		if err != nil {
			fmt.Println("Please enter a valid file.")
			return
		}else{
			var data []string
			dat := string(da)
			data = car.StringToArray(dat)
			for i := 0; i < len(data); i++ {
				fmt.Println(data[i])
			}
		}

	}else {
		fmt.Print("Incorrect argument format! Please try again.")
	}
	
}

