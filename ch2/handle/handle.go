package handle

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	source := map[string]string{
		"Hello": "World",
	}
	jsBytes, err := json.Marshal(source)
	if err != nil {
		log.Print(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	gzWriter := gzip.NewWriter(w)
	defer gzWriter.Close()
	multiWriter := io.MultiWriter(os.Stdout, gzWriter)
	_, err = multiWriter.Write(jsBytes)
	if err != nil {
		log.Print(err.Error())
		return
	}
}
