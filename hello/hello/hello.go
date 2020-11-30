package hello

import (
	"bufio"
	"io"

	"github.com/pkg/errors"
)

func Greet(writer io.Writer) error {
	buf := bufio.NewWriter(writer)
	_, err := buf.WriteString("Hello,World!\n")
	if err != nil {
		return errors.Wrap(err, "failed to write string")
	}
	if err := buf.Flush(); err != nil {
		return errors.Wrap(err, "failed to flush")
	}
	return nil
}
