/* vim: set ts=4 sw=4: */

/*
Response

Response はgatewayの実装
*/
package gateway

import (
	"bytes"
)

type IResponse interface {
	Buffer() *bytes.Buffer
}

type response struct {
	stringBuffer *bytes.Buffer
}

func NewResponse() IResponse {
	return &response{new(bytes.Buffer)}
}

func (r *response) Buffer() *bytes.Buffer {
	return r.stringBuffer
}
