package main

import (
	"reloaded"
	"fmt"
	"io"
	"os"
)
//There is an error in resulte file. It gives a space in front of the word. fix it.
func main() {
	files := os.Args[1:]
	if len(files) == 2 {
		dataFile := files[0]
		dat, err := os.ReadFile(dataFile)
		if err != nil {
			fmt.Println("Please enter a valid file.")
			return
		}
		data := reloaded.StringToArray(string(dat))
		data = reloaded.Check(data)
		data = reloaded.VowelLetter(data)
		newData := reloaded.ArrayToString(data)

		resFile, err := os.Create(files[1])
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		io.WriteString(resFile, newData)
	} else {
		fmt.Print("Incorrect argument format! Please try again.")
	}
}
