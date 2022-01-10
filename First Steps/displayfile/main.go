package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {

		file, err := os.Open("quest8.txt")
		if err != nil {
			fmt.Printf("the mistake is : %v\n", err.Error())
		}

		arr := make([]byte, 15)

		file.Read(arr)

		fmt.Print(string(arr))

		file.Close()
	}
}
