package main

import (
	"fmt"

	"github.com/SGWinter/internal/config"
)

func main() {
	testFile := config.Read()
	fmt.Printf("%v", testFile)
}
