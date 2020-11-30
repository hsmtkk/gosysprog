package main

import (
	"log"
	"os"

	"github.com/hsmtkk/gosysprog/ch1/hello"
)

func main() {
	if err := hello.Greet(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
