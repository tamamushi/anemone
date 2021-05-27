/* vim: set ts=4 sw=4: */

package codes

type Code int

const (
	OK Code = iota + 1
	NotEnoughArgument
	ForTestCode
)

func (c Code) String() string {
	return [...]string{"OK", "NotEnoughArgument", "ForTestCode"}[c-1]
}
