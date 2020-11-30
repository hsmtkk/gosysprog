package copy_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/hsmtkk/gosysprog/ch3/copy"
	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	testDir := t.TempDir()
	dst := filepath.Join(testDir, "sample.txt")
	src := "./sample.txt"
	err := copy.Copy(dst, src)
	assert.Nil(t, err)
	orig, err := ioutil.ReadFile(dst)
	assert.Nil(t, err)
	copied, err := ioutil.ReadFile(src)
	assert.Nil(t, err)
	assert.Equal(t, orig, copied)
}
