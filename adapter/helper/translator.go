/* vim: set ts=4 sw=4: */

package helper

type Translator interface {
}

type translation struct {
}

func NewTranslator() translation {
	return &translation{}
}
