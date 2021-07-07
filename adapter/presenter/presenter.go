/* vim: set ts=4 sw=4: */

package presenter

import (
//"anemone/adapter/gateway"
)

type Presenter interface {
	Command() []string
	Arguments() []string
}
