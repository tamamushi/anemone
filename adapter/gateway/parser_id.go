/* vim: set ts=4 sw=4: */

/*
Gateway はgatewayである
*/
package gateway

import (
	"encoding/json"
)

type UserIdFormatParser struct {
}

func (p *UserIdFormat) TryParse(data string, m interface{}) error {
	if err := json.Unmarshal([]byte(data), m); err != nil {
		return err
	}
	return nil
}
