package main

import (
	"fmt"

	"github.com/SGWinter/gator/internal/config"
)

func main() {
	file, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	file.SetUser("sean")
	file, err = config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v, %v\n", file.URL, file.USER)
}
