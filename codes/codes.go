/* vim: set ts=4 sw=4: */

package codes

type Code int

const (
	OK Code = iota + 1
	NotEnoughArgument
	Unknown
	ForTestCode
)

func (c Code) String() string {
	return [...]string{"OK", "NotEnoughArgument", "Unknown", "ForTestCode"}[c-1]
}
