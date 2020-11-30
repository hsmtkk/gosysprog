package copy

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

func Copy(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, "failed to open source")
	}
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to open destination")
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		return errors.Wrap(err, "failed to copy")
	}
	return nil
}
