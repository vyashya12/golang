package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("Generated ID: ", id)
}

//go mod tidy
// we can run the above to clean up our go file

// go build "file name"
// converts to binary

// to run the binary just go to the folder with the file and ./filenname
