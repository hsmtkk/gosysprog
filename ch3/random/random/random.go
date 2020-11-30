package random

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

func WriteRandomBytes(size int64, writer io.Writer) error {
	copied, err := io.CopyN(writer, rand.Reader, size)
	if err != nil {
		return errors.Wrap(err, "failed to copy")
	}
	if size != copied {
		return fmt.Errorf("want %d bytes, got %d bytes", size, copied)
	}
	return nil
}
