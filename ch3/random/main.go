package main

import (
	"log"
	"os"
	"strconv"

	"github.com/hsmtkk/gosysprog/ch3/random/random"
)

func main() {
	path := os.Args[1]
	size, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	writer, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()
	if err := random.WriteRandomBytes(size, writer); err != nil {
		log.Fatal(err)
	}
}
