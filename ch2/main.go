package main

import (
	"net/http"

	"github.com/hsmtkk/gosysprog/ch2/handle"
)

func main() {
	http.HandleFunc("/", handle.MyHandler)
	http.ListenAndServe(":8080", nil)
}
