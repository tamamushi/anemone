/* vim: set ts=4 sw=4: */

package codes

type Code string

const (
	OK                  Code = "OK"
	NotEnoughArgument   Code = "NotEnoughArgument"
	InvalidArgument     Code = "InvalidArgument"
	InternalServerError Code = "InternalServerError"
	UnSupportedMethod   Code = "UnSupportedMethod"
	Unknown             Code = "Unknown"
	Nil                 Code = "Nil"
	ForTestCode         Code = "ForTestCode"
)

func (c Code) String() string {
	return string(c)
}
