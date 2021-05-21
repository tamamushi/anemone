/* vim: set ts=4 sw=4: */

package gateway

import (
	"bytes"
	"io"
)

type IResponse interface {
	Buffer() io.Writer
}

type response struct {
	writer io.Writer
}

func NewResponse() IResponse {
	return &response{new(bytes.Buffer)}
}

func (r *response) Buffer() io.Writer {
	return r.writer
}
