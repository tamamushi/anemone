/* vim: set ts=4 sw=4: */

package errors

import (
	"fmt"
	"golang.org/x/xerrors"

	"anemone/codes"
)

type Errors interface {
	Error() string
}

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

func Newf(c codes.Code, msg string, a ...interface{}) error {
	return &structError{code: c, err: xerrors.New(fmt.Sprintf(msg, a...))}
}

func Messagef(msg string, a ...interface{}) string {
	return fmt.Sprintf(msg, a...)
}

func Code(err error) codes.Code {
	if e, ok := err.(*structError); ok {
		return e.code
	} else if err == nil {
		return codes.Nil
	}
	return codes.Unknown
}
