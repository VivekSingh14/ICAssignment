package server

import (
	"log"
	"os"
)

func MakeFile() {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

}
