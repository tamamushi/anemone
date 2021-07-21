/* vim: set ts=4 sw=4: */

/*
Gateway はgatewayである
*/
package gateway

import (
	"encoding/json"
)

type IParser interface {
	SetParser(interface{})
}

type IFormatParser interface {
	TryParse(string) error
}

type Input struct {
}

type Output struct {
}
