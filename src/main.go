package main

import (
	"blip-fullstack.com/test/src/parsers"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("commits.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	commits, errCSV := parsers.ParseCSV(file)
	if errCSV != nil {
		fmt.Println("Error reading CSV:", errCSV)
		return
	}

	fmt.Printf("%+v\n", commits)
}
