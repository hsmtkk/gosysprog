package handle_test

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsmtkk/gosysprog/ch2/handle"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestMyHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handle.MyHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	assert.Nil(t, err)
	defer resp.Body.Close()
	gzCompressed, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	want := []byte(`{"Hello":"World"}`)
	got, err := decompress(gzCompressed)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func decompress(compressed []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, errors.Wrap(err, "failed to recognize as gzip stream")
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, reader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to copy bytes")
	}
	return buf.Bytes(), nil
}
