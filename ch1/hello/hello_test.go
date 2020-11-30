package hello_test

import (
	"bytes"
	"testing"

	"github.com/hsmtkk/gosysprog/hello/hello"
	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	var buf bytes.Buffer
	err := hello.Greet(&buf)
	assert.Nil(t, err)
	want := []byte("Hello,World!\n")
	got := buf.Bytes()
	assert.Equal(t, want, got)
}
