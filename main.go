package main

import (
	"fmt"

	"github.com/SGWinter/gator/internal/config"
)

func main() {
	testFile, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", testFile.URL)
}
