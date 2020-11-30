package random_test

import (
	"bytes"
	"testing"

	"github.com/hsmtkk/gosysprog/ch3/random/random"
	"github.com/stretchr/testify/assert"
)

func TestWriteRandomBytes(t *testing.T) {
	var buf bytes.Buffer
	err := random.WriteRandomBytes(1024, &buf)
	assert.Nil(t, err)
	assert.NotNil(t, buf.Bytes())
}
