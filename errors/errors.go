/* vim: set ts=4 sw=4: */

package errors

import (
	"fmt"
	"golang.org/x/xerrors"

	"anemone/codes"
)

type AnemoneError interface {
	Code() codes.Code
	Error() string
}

type structError struct {
	code codes.Code
	err  error
}

func (e structError) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s", e.code, e.err)
}

func New(c codes.Code, msg string) AnemoneError {
	return structError{code: c, err: xerrors.New(msg)}
}

func (e structError) Code() codes.Code {
	return e.code
}
