/* vim: set ts=4 sw=4: */

package errors

import (
	"fmt"
	"golang.org/x/xerrors"

	"anemone/codes"
)

type structError struct {
	code codes.Code
	err  error
}

func (e *structError) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s", e.code, e.err)
}

func New(c codes.Code, msg string) error {
	return &structError{code: c, err: xerrors.New(msg)}
}

func Code(err error) codes.Code {
	if e, ok := err.(*structError); ok {
		return e.code
	}
	return codes.Unknown
}
